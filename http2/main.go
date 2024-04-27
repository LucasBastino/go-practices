package main

import (
	"fmt"
	"net/http"

	"github.com/LucasBastino/practicas-go/http2/api"
	"github.com/gorilla/mux"
)

func main() {
	// instancio un enrutador
	r := mux.NewRouter()

	// instancio la api controladora
	a := api.Controller{}

	// inicio le paso el enrutador e inicio el enrutamiento
	a.RegisterRoutes(r)

	svr := http.Server{
		Addr:    ":8085",
		Handler: r,
	}

	fmt.Println("Listening on port", svr.Addr)
	svr.ListenAndServe()
}

// func getUsers(path string) []User {
// 	file, _ := os.Open(path)
// 	var users []User
// 	json.NewDecoder(file).Decode(&users)
// 	return users
// }
