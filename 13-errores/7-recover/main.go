package main

import "fmt"

func main() {
	a()
	fmt.Println("Imprimiendo desde Main")
}

func a() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperando i desde recover en a:", r)
		}
	}()
	b(0)
	fmt.Println("Retornando a main desde b") // esto no se va a ejecutar porque
	// despues del panic va directo al recover
}

func b(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
		// no se por que no se escribe asi y listo:
		// panic(i)

	} else {
		defer fmt.Println("Imprimiendo i en defer desde b:", i)
		fmt.Println("Imprimiendo i normal desde b:", i)
		b(i + 1)
	}
}

// cuando paniquea, va directamente al recover dandole el valor que pasamos por parametro
