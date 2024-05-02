package api

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func (c *Controller) editUser(w http.ResponseWriter, r *http.Request) {
	var userToEdit models.User
	users := c.decodeUsers()
	// inserto las variables del json pasado por el body en la variable userToEdit
	err := json.NewDecoder(r.Body).Decode(&userToEdit)
	if err != nil {
		fmt.Println("error decoding body in editUser")
		panic(err.Error())
	}
	// recojo las variables escritas en la URL
	URLParams := mux.Vars(r)
	// en este caso solo me interesa el id
	id, err := strconv.Atoi(URLParams["id"])
	if err != nil {
		fmt.Println("error converting idParam to int")
	}

	users, validation := func() ([]models.User, bool) {
		for index, user := range users {
			if user.Id == id {
				// si coincide el param id con el id del user borra el user anterior y mete el modificado
				usersTemp := append(users[:index], userToEdit)
				users = append(usersTemp, users[index+1:]...)
				return users, true
			}
		}

		return users, false
	}()

	if validation {
		c.saveUsers(users)
	} else {
		w.WriteHeader(http.StatusNotModified)
	}

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
