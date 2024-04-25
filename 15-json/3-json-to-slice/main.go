package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	UserId    int // no hace falta ponerles el `json: userId` en minuscula porque el Unmarshal te lo toma igual
	Id        int
	Title     string
	Completed bool
}

// func (u User) String() string {
// 	return fmt.Sprintf("id: %d, title: %s", u.Id, u.Title)
// }

func main() {
	dataJson, err := os.ReadFile("./data.json")
	if err != nil {
		fmt.Print("error reading file")
		panic(err.Error())
	}
	// dataString := string(dataJson)
	// //  si no lo pasas a string salen todos bytes
	// fmt.Println(dataString)
	// igualmente no necesito convertirlo, json.Unmarshal me pide como primer parametro un []byte
	// users := []User{} -> es lo mismo
	var users []User
	err = json.Unmarshal(dataJson, &users)
	// si los atributos comienzan con mayuscula Unmarshal los asigna igual aunque en el .json empiecen con minuscula
	if err != nil {
		fmt.Println("error unmarshaling data")
		panic(err.Error())
	}
	// fmt.Println(users)

	for _, user := range users {
		fmt.Println(user)
	}
}
