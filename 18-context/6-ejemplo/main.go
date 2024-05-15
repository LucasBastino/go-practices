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
	// ctx := context.Background()
	// ctx, cancelTimeout := context.WithTimeout(ctx, 2*time.Millisecond)
	// defer cancelTimeout()
	// contextWithtimeout no funciona adentro de handlers, averiguar por que

	// este cancel esté o no es lo mismo, ya viene en el r.Context()
	// ctx, cancel := context.WithCancel(ctx)
	// defer cancel()

	select {
	case <-ctx.Done():
		fmt.Println("contexto terminado, sea por timeout o por cancel")
	case <-time.After(2 * time.Second):
		fmt.Println("pasaron 2 segundos")
	}

}
