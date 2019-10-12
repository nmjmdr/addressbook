package router

import (
	"addressbook/api/status"
	"github.com/gorilla/mux"
	"net/http"
)

// Start starts the router and add routes
func Start(listenAddress string) {
	r := mux.NewRouter()
	r.HandleFunc("/api/status", status.Handle).Methods("GET")

	http.ListenAndServe(listenAddress, r)
}
