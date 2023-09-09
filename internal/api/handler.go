package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type APIHandler struct {
}

func NewAPIHandler() *APIHandler {
	return &APIHandler{}
}

func optimizerHandler(r *mux.Router) {
	oh := NewAPIHandler()

	r.HandleFunc("/inefficient", oh.getInefficientHosts).Methods("GET")
}

// TODO: Add logic for inefficient hosts
func (ah *APIHandler) getInefficientHosts(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}
