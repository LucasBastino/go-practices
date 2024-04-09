package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Error escribiendo en un archivo con log.SetOutput")
	f, err := os.Create("log-error.txt")
	if err != nil {
		log.Fatalln("Ha ocurrido un error:", err)
	}
	// obliga al archivo a cerrarse antes de que termine de ejecutarse la funcion
	defer f.Close()

	// indica que el log del error se guarde en el archivo "f"
	log.SetOutput(f)

	f2, err := os.Open("archivoFicticio.txt")
	if err != nil {
		log.Println("Ha ocurrido un error:", err)
	}
	f2.Close()

	fmt.Println("Revisa el archivo log-error.txt")
}

// fmt.Printnln(error) -> muestra el error
// log.Println(error) -> muestra el error con la hora y fecha
// log.Panicln(error) -> funciones defer corren y se puede hacer recover
// panic(error)
