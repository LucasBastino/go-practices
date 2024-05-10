package main

import (
	"fmt"
	"net/http"

	"github.com/LucasBastino/practicas-go/16-http/v1/api"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	a := &api.API{}
	a.RegisterRoutes(r)

	srv := &http.Server{
		Addr:    ":8085",
		Handler: r,
	}
	fmt.Println("Listening on port", srv.Addr)
	srv.ListenAndServe()

}

// func HandleHome(users []User) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		encoder := json.NewEncoder(w)
// 		encoder.Encode(users)
// 	}
// }
