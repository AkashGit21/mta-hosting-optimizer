package api

import (
	"net/http"

	"github.com/AkashGit21/mta-hosting-optimizer/internal/utilities"
	"github.com/gorilla/mux"
)

func New() (*mux.Router, error) {
	router := mux.NewRouter()

	hostsRouter := router.PathPrefix("/hosts").Subrouter()
	hostsRouter.Use(PanicRecoveryMiddleware)
	optimizerHandler(hostsRouter)

	return router, nil
}

func PanicRecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				// Handle the panic
				utilities.InfoLog("Panic recovered: ", r)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
