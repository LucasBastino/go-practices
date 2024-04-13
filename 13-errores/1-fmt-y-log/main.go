package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("archivoFicticio.txt")
	if err != nil {
		fmt.Println("Error con fmt.Println")
		fmt.Println("Ha ocurrido un error:", err)
		fmt.Println("Error con log.Println")
		log.Println("Ha ocurrido un error:", err)
	}
}

// fmt.Println no devuelve el error, solo lo imprime
