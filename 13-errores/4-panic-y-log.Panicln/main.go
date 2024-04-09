package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Impresion de mensaje diferida")
	_, err := os.Open("archivo-ficticio.txt")
	if err != nil {
		// log.Panicln(err)
		panic(err)
	}
}

// Con panic se puede usar recover y ademas las defer corren normalmente
