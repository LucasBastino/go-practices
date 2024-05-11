package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancelTimeout := context.WithTimeout(ctx, time.Second*5)
	defer cancelTimeout()

	// printMessage(ctx, time.Second*2, "hola")
	printMessage(ctx, time.Second*7, "hola")
}

func printMessage(ctx context.Context, t time.Duration, message string) {
	select {
	case <-ctx.Done():
		fmt.Println("timeout, 5 seconds transcurred")
	// despues del tiempo t, time.After devuelve un struct vacio, o sea una señal
	case <-time.After(t):
		fmt.Println("no timeout, message:", message)
	}
}
