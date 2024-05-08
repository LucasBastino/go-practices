package api

import (
	"github.com/gorilla/mux"
)

func (api *API) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("GET /", api.HandleHome)
}
