package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	users := []string{"fran moreno", "gustavinho", "todopasa"}
	ctx := context.Background()
	ctx, cancelTimeout := context.WithTimeout(ctx, time.Second*5)
	defer cancelTimeout()

	printUsers(ctx, users)

}

func printUsers(ctx context.Context, users []string) {
	for i := range users {
		select {
		case <-ctx.Done():
			fmt.Println("timeout, context done, 5 seconds transcurred")
			return
		case <-time.After(time.Second * 2):
			fmt.Println(users[i])
		}
	}
}
