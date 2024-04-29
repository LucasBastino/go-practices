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

func (c *Controller) getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	id, err := strconv.Atoi(idParam)

	if err != nil {
		fmt.Println("error converting idParam to id")
		w.WriteHeader(http.StatusBadRequest)
	}

	users := c.decodeUsers()
	var user models.User
	// si los usuarios estan ordenandos consecutivamente por id
	// user := users[id-1]

	// y si no estan ordenados consecutivamentes por id
	for index, u := range users {
		if u.Id == id {
			user = users[index]
			json.NewEncoder(w).Encode(user)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func (c *Controller) createUser(w http.ResponseWriter, r *http.Request) {
	users := c.decodeUsers()
	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Printf("%+v\n", user)
	users = append(users, user)
	c.setUsersPrueba(user)

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

var users2 []models.User

func (c *Controller) setUsersPrueba(user models.User) {
	users2 = c.decodeUsers()
	users2 = append(users2, user)
}

func (c *Controller) getUsersPrueba(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users2)
}
func (c *Controller) getUserPrueba(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	id, _ := strconv.Atoi(idParam)

	// users := c.decodeUsers()
	var user models.User

	// y si no estan ordenados consecutivamentes por id
	for index, u := range users2 {
		if u.Id == id {
			user = users2[index]
			json.NewEncoder(w).Encode(user)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}
