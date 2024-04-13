package main

import (
	"errors"
	"fmt"
)

func main() {
	raiz, err := sqrt(-2)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Raiz:", raiz)
	}

}

func sqrt(num int) (int, error) {
	if num < 0 {
		error1 := errors.New("No existe raiz cuadrada de un numero negativo")
		return 0, error1
		// tambien puede ser asi
		// return 0, errors.New("No existe raiz cuadrada de un numero negativo")
		// return 0, fmt.Errorf("No existe raiz cuadrada de un numero negativo: %v", num)
	} else {
		return 40, nil
	}
}

// errors.New() devuelve el error
