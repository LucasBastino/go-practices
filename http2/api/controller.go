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
	} else if from < 1 || to < 1 || from > to {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		validation, idFrom, idTo := c.validationUsersFromTo(users, from, to)
		if validation {
			// para mandar los datos de la variable no hace falta el puntero
			// desde(inclusive) hasta(inclusive)
			json.NewEncoder(w).Encode(users[idFrom : idTo+1])
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}

}

func (c *Controller) validationUsersFromTo(users []models.User, from, to int) (bool, int, int) {
	var idFrom int
	var idTo int
	var idFromFounded bool
	var idToFounded bool
	for index, user := range users {
		if user.Id == from {
			idFrom = index
			idFromFounded = true
			break
		} else {
			idFromFounded = false
		}
	}

	for index, user := range users {
		if user.Id == to {
			idTo = index
			idToFounded = true
			break
		} else {
			idToFounded = false
		}
	}
	fmt.Println(idFrom, idTo)
	fmt.Println(idFromFounded, idToFounded)
	if idFromFounded && idToFounded { // == true
		return true, idFrom, idTo
	} else {
		return false, 0, 0
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
	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Printf("%+v\n", user)
	users := c.decodeUsers()
	users = append(users, user)
	c.saveUsers(users)
	w.WriteHeader(http.StatusCreated)
}

func (c *Controller) deleteUser(w http.ResponseWriter, r *http.Request) {
	userParams := mux.Vars(r)
	idParam := userParams["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		fmt.Println("err converting idParam to int")
		panic(err.Error())
	}

	users := c.decodeUsers()
	for index, user := range users {
		if user.Id == id {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	c.saveUsers(users)
}

func (c *Controller) saveUsers(users []models.User) {
	usersJson, err := json.Marshal(users)
	if err != nil {
		fmt.Println("error marshaling users")
		panic(err.Error())
	}
	os.WriteFile(c.getPath(), usersJson, 0644)
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
	defer file.Close()

	var users []models.User
	// para guardar los datos en la variable se requiere de un puntero
	err = json.NewDecoder(file).Decode(&users)
	if err != nil {
		fmt.Println("error decoding users")
		panic(err.Error())
	}
	return users
}
