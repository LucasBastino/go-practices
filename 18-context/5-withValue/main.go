package main

import (
	"context"
	"fmt"
)

// esto en realidad sirve mas para cosas efimeras como pasar un value en una req
// no es para pasar values como argumentos de una funcion, pero aca se ejemplifica de esa forma
func main() {
	users := []string{"lucas", "leandro"}
	ctx := context.Background()
	ctx = context.WithValue(ctx, "name", "lucas") // arreglar esto

	test(ctx, users)
}

func test(ctx context.Context, users []string) {
	name, ok := ctx.Value("name").(string)
	if !ok {
		fmt.Println("key not found")
	} else {
		for _, user := range users {
			if user == name {
				fmt.Println("the user", name, "exists")
				return
			}
		}
		fmt.Println("not coincidence found")
	}

}
