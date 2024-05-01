package api

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/LucasBastino/practicas-go/http2/models"
)

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

func (c *Controller) saveUsers(users []models.User) {
	usersJson, err := json.Marshal(users)
	if err != nil {
		fmt.Println("error marshaling users")
		panic(err.Error())
	}
	os.WriteFile(c.getPath(), usersJson, 0644)
}
