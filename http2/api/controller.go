package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/LucasBastino/practicas-go/http2/models"
	"github.com/gorilla/schema"
)

// controllers

type Controller struct{}

func (c *Controller) renderHome(w http.ResponseWriter, r *http.Request) {
	str := "este es el home index"
	json.NewEncoder(w).Encode(str)
}

func (c *Controller) getUsers(w http.ResponseWriter, r *http.Request) {
	users := c.decodeUsers()
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

func (c *Controller) getUser(w http.ResponseWriter, r *http.Request) {
	// idParam := r.URL.Query().Get("id")

	// params := mux.Vars(r)
	// idParam := params["id"]
	// id, err := strconv.Atoi(idParam)

	// if err != nil {
	// 	fmt.Println("error converting idParam to id")
	// 	w.WriteHeader(http.StatusBadRequest)
	// }

	userParams := models.UserParams{}
	schema.NewDecoder().Decode(userParams, r.URL.Query())
	users := c.decodeUsers()
	user := users[userParams.Id]
	fmt.Println(user)
	json.NewEncoder(w).Encode(user)
}

func (c *Controller) getPath() string {
	return "./data.json"
}

func (c *Controller) decodeUsers() []models.User {
	path := c.getPath()
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("error opening file")
		panic(err.Error())
	}
	var users []models.User
	// para guardar los datos en la variable se requiere de un puntero
	err = json.NewDecoder(file).Decode(&users)
	if err != nil {
		fmt.Println("error decoding users")
		panic(err.Error())
	}
	return users
}
