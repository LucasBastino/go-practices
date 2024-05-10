package api

import (
	"net/http"
)

// enrutador

func (c *Controller) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("GET /", c.renderHome)
	r.HandleFunc("GET /users", c.getUsers)
	r.HandleFunc("GET /users/", c.getUsers)
	r.HandleFunc("POST /users/create", c.createUser)
	r.HandleFunc("PUT /users/edit/{id}", c.editUser)
	r.HandleFunc("GET /users/{id}", c.getUser)
	r.HandleFunc("DELETE /users/delete/{id}", c.deleteUser)
}
