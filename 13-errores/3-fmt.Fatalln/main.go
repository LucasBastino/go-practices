package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer fmt.Println("Impresion de mensaje diferida")
	_, err := os.Open("archivo-ficticio.txt")
	if err != nil {
		log.Fatalln(err)
	}
}

// log.Fatalln(error) -> hace kill de todos los procesos, llama a os.Exit(), no ejecuta las defer
