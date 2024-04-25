package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type User struct {
	UserId    int
	Id        int
	Title     string
	Completed bool
}

func (u User) String() string {
	return fmt.Sprintf("id %v, title %s", u.Id, u.Title)
}

func main() {
	// decodifico el data.json y lo guardo en users
	file, err := os.Open("./data.json")
	if err != nil {
		fmt.Println("error opening data.json")
		panic(err.Error())
	}
	var users []User
	json.NewDecoder(file).Decode(&users)

	http.HandleFunc("/", handler(users))
	http.ListenAndServe(":8051", nil)

}

func handler(users []User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// con el .Encode(users) le estas mandando la info
		err := json.NewEncoder(w).Encode(users)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}
