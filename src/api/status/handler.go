package status

import (
	"fmt"
	"net/http"
)

// Handle - Status Handler
func Handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}
