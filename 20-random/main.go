package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	for range 10 {

		fmt.Println(rand.IntN(25000000) + 20000000)
	}
}
