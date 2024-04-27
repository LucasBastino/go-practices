package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// enrutador

func (c *Controller) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/", c.renderHome).Methods(http.MethodGet)
	r.HandleFunc("/users", c.sendUsers).Methods(http.MethodGet)
}
