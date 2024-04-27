package api

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/LucasBastino/practicas-go/http2/models"
)

// controllers

type Controller struct{}

func (c *Controller) renderHome(w http.ResponseWriter, r *http.Request) {
	str := "este es el home index"
	json.NewEncoder(w).Encode(str)
}

func (c *Controller) sendUsers(w http.ResponseWriter, r *http.Request) {
	path := c.getPath()
	file, _ := os.Open(path)

	var users []models.User
	// para guardar los datos en la variable se requiere de un puntero
	json.NewDecoder(file).Decode(&users)

	limitParam := r.URL.Query().Get("limit")
	limit, _ := strconv.Atoi(limitParam)

	if limit < 0 || limit > 200 {
		w.WriteHeader(http.StatusBadRequest)
	} else if limit == 0 {
		json.NewEncoder(w).Encode(users)
	} else {
		// para mandar los datos de la variable no hace falta el puntero
		json.NewEncoder(w).Encode(users[:limit])
	}
}

func (c *Controller) getPath() string {
	return "./data.json"
}
