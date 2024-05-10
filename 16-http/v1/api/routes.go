package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (api *API) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/", api.HandleHome).Methods(http.MethodGet)
}
