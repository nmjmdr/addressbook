package autopilot

import (
	"addressbook/models"
	"net/http"
	"github.com/pkg/errors"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"bytes"
)

/*
type AutopilotProxy interface {
	Get(idOrEmail string) ([]models.Contact, error)
	Upsert(contact models.Contact) error
}
*/

type Proxy struct {
	baseUrl string
	apiKey string
}

const APIHeaderKey = "autopilotapikey"

func parseContact(body []byte) ([]models.Contact, error) {
	contact := models.Contact{}
	err := json.Unmarshal(body, &contact)
	if err != nil {
		return []models.Contact{}, err
	}
	return []models.Contact {
		contact,
	}, nil
}

type ErrorResponse struct {
	Error string `json:"error"`
	Message string `json:"message"`
}

func parseErrorResponse(body []byte) (ErrorResponse, error) {
	errResponse := ErrorResponse{}
	err := json.Unmarshal(body, &errResponse)
	if err != nil {
		return ErrorResponse{}, err
	}
	return errResponse, nil
}

func (p *Proxy) Get(idOrEmail string) ([]models.Contact, error) {
	url := fmt.Sprintf("%s/%s",p.baseUrl,idOrEmail)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []models.Contact{}, err
	}
	req.Header.Add(APIHeaderKey, p.apiKey)
	resp, err := client.Do(req)
	if err != nil {
		return []models.Contact{}, errors.Wrap(err,"Unable to fetch contact from API")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []models.Contact{}, errors.Wrap(err,"Unable to fetch contact from API")
	}

	if resp.StatusCode == http.StatusOK {
		return parseContact(body)
	} 

	if resp.StatusCode == http.StatusNotFound {
		return []models.Contact{}, nil
	}

	errResponse, err := parseErrorResponse(body)
	if err != nil {
		return []models.Contact{}, errors.Wrap(err, "Unable to read response body")
	}
	return []models.Contact{}, errors.New(errResponse.Message)
}

type UpsertRequestBody struct {
	Contact models.Contact `json:"contact"`
}

func (p *Proxy) Upsert(contact models.Contact) error {
	url := p.baseUrl
	client := &http.Client{}

	jsonBytes, _ := json.Marshal(UpsertRequestBody{ Contact: contact })
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	req.Header.Add(APIHeaderKey, p.apiKey)
    req.Header.Add("Content-Type", "application/json")
	
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err,"Unable to invoke API to upsert contact")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err,"Unable to invoke API to upsert contact")
	}

	if resp.StatusCode == http.StatusOK {
		return nil
	} 

	errResponse, err := parseErrorResponse(body)
	if err != nil {
		return errors.Wrap(err, "Unable to read response body")
	}
	return  errors.New(errResponse.Message)
}

func NewAutoPilotProxy(baseUrl string, apiKey string) *Proxy {
	return &Proxy{
		baseUrl: baseUrl,
		apiKey: apiKey,
	}
}