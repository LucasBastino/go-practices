package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctxTimeout, cancelTimeout := context.WithTimeout(ctx, time.Second*5)
	defer cancelTimeout()
	ctxCancel, cancel := context.WithCancel(ctx)
	time.Sleep(time.Second * 2)
	// defer cancel()
	cancel()
	test(ctxTimeout, ctxCancel)
}

func test(ctxTimeout context.Context, ctxCancel context.Context) {
	// users := []string{"funebrero96", "fran moreno", "gustavinho"}
	// ch := make(chan string)
	// for _, user := range users {
	// 	ch <- user
	// }
	select {
	// ctx.Done() devuelve un chan de struct vacio, envia una señal
	case <-ctxTimeout.Done():
		fmt.Println("context timeout done, 5 seconds transcurred")
	case <-ctxCancel.Done():
		fmt.Println("context cancelled")
	}

}
