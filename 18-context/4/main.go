package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	svr := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(handler),
	}
	svr.ListenAndServe()
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	select {
	case <-ctx.Done():
		fmt.Println("request is cancelled, context is done")
	case <-time.After(time.Second * 3):
		fmt.Println("message printed after 3 seconds")
	}
}
