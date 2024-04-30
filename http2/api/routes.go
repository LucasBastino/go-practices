package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// enrutador

func (c *Controller) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/", c.renderHome).Methods(http.MethodGet)
	r.HandleFunc("/users", c.getUsers).Methods(http.MethodGet)
	r.HandleFunc("/users/", c.getUsers).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", c.getUser).Methods(http.MethodGet)
	r.HandleFunc("/users/create", c.createUser).Methods(http.MethodPost)
	r.HandleFunc("/users/delete/{id}", c.deleteUser).Methods(http.MethodDelete)
}
