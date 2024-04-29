package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// enrutador

func (c *Controller) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/", c.renderHome).Methods(http.MethodGet)
	r.HandleFunc("/users", c.getUsers).Methods(http.MethodGet)
	r.HandleFunc("/users/", c.getUsersPrueba).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", c.getUserPrueba).Methods(http.MethodGet)
	r.HandleFunc("/createUser", c.createUser).Methods(http.MethodPost)
}
