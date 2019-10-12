package router

import (
	"addressbook/api/status"
	"github.com/gorilla/mux"
	"net/http"
	"addressbook/store"
	"addressbook/api/contacts"
)

// Start starts the router and add routes
func Start(listenAddress string,st store.Store) {
	r := mux.NewRouter()
	r.HandleFunc("/api/status", status.Handle).Methods("GET")

	contactsHandler := contacts.NewHandler(st)	

	r.HandleFunc("/api/contacts/{idOrEmail}", contactsHandler.Get).Methods("GET")

	r.HandleFunc("/api/contacts", contactsHandler.Upsert).Methods("POST")

	r.HandleFunc("/api/contacts", contactsHandler.Upsert).Methods("PUT")

	http.ListenAndServe(listenAddress, r)
}
