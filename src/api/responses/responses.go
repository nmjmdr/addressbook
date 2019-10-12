package responses

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	w http.ResponseWriter
}

func NewResponse(w http.ResponseWriter) *Response {
	return &Response{
		w: w,
	}
}

type ErrorResponse struct {
	Err string `json:"error"`
}

func (r *Response) InternalError(err error) {
	r.sendError(err.Error(), http.StatusInternalServerError)
}

func (r *Response) BadRequest(err error) {
	r.sendError(err.Error(), http.StatusBadRequest)
}

func (r *Response) sendError(err string, statusCode int) {
	r.w.WriteHeader(statusCode)
	r.w.Header().Set("Content-Type", "application/json")
	js, _ := json.Marshal(ErrorResponse{
		Err: err,
	})
	r.w.Write(js)
}

func (r *Response) NotFound(message string) {
	r.sendError(message, http.StatusNotFound)
}

func (r *Response) Json(data []byte) {
	r.w.WriteHeader(http.StatusOK)
	r.w.Header().Set("Content-Type", "application/json")
	r.w.Write(data)
}

func (r *Response) OK() {
	r.w.WriteHeader(http.StatusOK)
}
