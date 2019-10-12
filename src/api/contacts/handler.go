package contacts

import (
	"fmt"
	"net/http"
	"addressbook/store"
	"github.com/gorilla/mux"
	"addressbook/api/responses"
	"encoding/json"
	"io/ioutil"
	"addressbook/models"
	"github.com/pkg/errors"
)

type Handler struct {
	store store.Store
}

func NewHandler(st store.Store) *Handler {
	return &Handler{
		store: st,
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idOrEmail := params["idOrEmail"]
	contacts, err := h.store.Get(idOrEmail)

	response := responses.NewResponse(w)

	if err != nil {
		response.InternalError(err)
		return
	}

	if len(contacts) == 0 {
		response.NotFound(fmt.Sprintf("Contact with id or email: %s not found",idOrEmail))
		return
	}

	js, _  := json.Marshal(contacts[0])
	response.Json(js)
}

func (h *Handler) Upsert(w http.ResponseWriter, r *http.Request) {
	
	response := responses.NewResponse(w)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.BadRequest(err)
		return
	}
	
	contact := models.Contact{}
	err = json.Unmarshal(body, &contact)
    if err != nil {
        response.BadRequest(errors.New("Input Error. Unable to parse request body as contact"))
		return
	}
	
	err = h.store.Upsert(contact)

	if err != nil {
		response.InternalError(err)
		return
	}
	response.OK()
}
