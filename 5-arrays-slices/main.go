package main

import "fmt"

func main() {
	fmt.Println("arrays and slices exercises:")
	arrays()
}

// una variable que contiene un vector se declara con
// var nombre [tamaño] tipo de dato
// var num1al10[10]int

func arrays() {
	// [5] indica la cantidad de elementos
	var arr1 [5]int
	arr2 := [5]int{1, 2, 3, 4, 5}
	// [...] indica que su tamaño va a ser igual a la cantidad de elementos
	arr3 := [...]int{2, 3, 4, 5, 6, 7, 8}
	// esto daria error, tiene mas elementos que su tamaño indicado
	// arrError := [3]int{1, 2, 3, 4, 5}

	// se pueden declarar matrices, es decir, arrays de multiples dimensiones
	var tableroDeAjedrez [8][8]int

	// 2 arrays de distinto tamaño son 2 arrays de distinto tipo
	// var [5]int y var[10]int son de distinto tipo
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(tableroDeAjedrez)

}
