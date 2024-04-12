package main

import (
	"fmt"
	"log"
)

func main() {
	raiz, err := sqrt(-2)
	if err != nil {
		log.Fatalln(err)
		// log.Println(err)

	} else {
		fmt.Println(raiz)
	}
}

type ubicacionError struct {
	altitud string
	latitud string
	err     error
}

func (uE ubicacionError) Error() string {
	return fmt.Sprintf("En la ubicacion %v, %v, ocurrio un error: %v\n", uE.altitud, uE.latitud, uE.err)
	// Sprintf RETORNA el string
}

func sqrt(num int) (int, error) {
	if num < 0 {
		error1 := fmt.Errorf("no existe raiz cuadrada de un numero negativo: %v", num)
		return 0, ubicacionError{"12.5", "42.17", error1}
	} else {
		return 40, nil
	}
}
