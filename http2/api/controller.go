package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/LucasBastino/practicas-go/http2/models"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

// controllers

type Controller struct{}

func (c *Controller) renderHome(w http.ResponseWriter, r *http.Request) {
	str := "este es el home index"
	json.NewEncoder(w).Encode(str)
}

func (c *Controller) getUsers(w http.ResponseWriter, r *http.Request) {
	// limitParam := r.URL.Query().Get("limit")
	// limit, _ := strconv.Atoi(limitParam)
	userParams := models.UserParams{}
	decoder := schema.NewDecoder()
	// el primer parametro del Decode hecho por un schema, debe ser un puntero a un struct
	err := decoder.Decode(&userParams, r.URL.Query())
	if err != nil {
		fmt.Println("error decoding URL")
		panic(err.Error())
	}
	users := c.decodeUsers()
	fmt.Println(userParams)
	from := userParams.From
	to := userParams.To
	if from == 0 && to == 0 {
		json.NewEncoder(w).Encode(users)
	} else if from < 1 || from > len(users) || to < 1 || to > len(users) || from > to {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		// para mandar los datos de la variable no hace falta el puntero
		// desde(inclusive) hasta(inclusive)
		json.NewEncoder(w).Encode(users[from-1 : to])
	}

}

func (c *Controller) postBook(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)

}

func (c *Controller) getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	id, err := strconv.Atoi(idParam)

	if err != nil {
		fmt.Println("error converting idParam to id")
		w.WriteHeader(http.StatusBadRequest)
	}

	users := c.decodeUsers()
	user :=
	user := users[id-1]
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
