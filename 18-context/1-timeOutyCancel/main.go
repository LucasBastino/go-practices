package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	ctxCancel, cancel := context.WithCancel(ctx)
	// defer cancel() // esto es para que se cierre y ahorrar recursos
	time.Sleep(time.Second * 2)
	cancel()

	ctxTimeout, cancelTimeout := context.WithTimeout(ctxCancel, time.Second*5)
	defer cancelTimeout() // esto es para que se cierre y ahorrar recursos
	test(ctxCancel, ctxTimeout)
}

func test(ctxCancel context.Context, ctxTimeout context.Context) {
	select {
	case <-ctxCancel.Done(): // ctx.Done() devuelve un struct vacio, o sea, una señal
		fmt.Println("context cancelled")
	case <-ctxTimeout.Done():
		fmt.Println("timeout, 5 seconds transcurred")
	}
}
