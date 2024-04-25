package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	UserId    int
	Id        int
	Title     string
	Completed bool
}

func main() {
	// os.Open devuelve un *File
	file, err := os.Open("./data.json")
	if err != nil {
		fmt.Println("error opening file")
		panic(err.Error())
	}
	// File tiene una funcion Read, por lo tanto implementa la interfaz Reader
	// json.NewDecoder pide como parametro un Reader, asi que se puede usar mi variable file tipo File
	decoder := json.NewDecoder(file)
	var users []User
	decoder.Decode(&users)

	for _, user := range users {
		fmt.Println(user)
	}
}
