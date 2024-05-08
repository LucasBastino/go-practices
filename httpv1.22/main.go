package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", renderHome)

	http.ListenAndServe("localhost:8085", mux)

}

func renderHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "holaaaa")
}
