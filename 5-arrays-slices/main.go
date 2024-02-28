package main

import "fmt"

func main() {
	fmt.Println("arrays and slices exercises:")
	// arrays()
	// appendSlice()
	// lenYcapSlice()
	// sliceConMake()
	// copySlice()
	// arrYsliceComoArgs()
	// forOptimizado()
	// viewsSlice()
	fmt.Println(funcionSumaConVariosArgs(10, 4, 5))

}

// una variable que contiene un vector se declara con
// var nombre [tamaño] tipo de dato
// var num1al10[10]int

// func arrays() {
// 	// [5] indica la cantidad de elementos
// 	var arr1 [5]int
// 	arr2 := [5]int{1, 2, 3, 4, 5}
// 	// [...] indica que su tamaño va a ser igual a la cantidad de elementos
// 	arr3 := [...]int{2, 3, 4, 5, 6, 7, 8}
// 	// esto daria error, tiene mas elementos que su tamaño indicado
// 	// arrError := [3]int{1, 2, 3, 4, 5}

// 	// se pueden declarar matrices, es decir, arrays de multiples dimensiones
// 	var tableroDeAjedrez [8][8]int

// 	// 2 arrays de distinto tamaño son 2 arrays de distinto tipo
// 	// var [5]int y var[10]int son de distinto tipo
// 	fmt.Println(arr1)
// 	fmt.Println(arr2)
// 	fmt.Println(arr3)
// 	fmt.Println(tableroDeAjedrez)

// }

// func appendSlice() {
// 	var slice1 []int
// 	slice2 := []int{1, 2}

// 	slice1 = append(slice2, 10, 20, 30)
// 	slice2 = append(slice1, 99, 88)

// 	fmt.Println(slice1)
// 	fmt.Println(slice2)

// }

// func lenYcapSlice() {
// 	var slice []int
// 	fmt.Println("Longitud:", len(slice), "Capacidad:", cap(slice))
// 	// aca el slice apunta a nil

// 	slice = append(slice, 1, 2, 3, 4)
// 	fmt.Println("Longitud:", len(slice), "Capacidad:", cap(slice))
// 	// se crea un slice con la longitud igual a la cantidad de elementos

// 	slice = append(slice, 5)
// 	fmt.Println("Longitud:", len(slice), "Capacidad:", cap(slice))
// 	// el slice ya llego a su capacidad maxima, por lo que crea otro con el doble de capacidad
// 	// le copia los valores del anterior slice y le agrega el elemento

// }

// func sliceConMake() {
// 	jugadoresDeUnEquipo := make([]string, 0, 11) (tipo, len, cap) la capacidad es opcional
// 	// se crea un slice vacio con capacidad maxima de 11,
// 	// no puede haber 12 jugadores en un equipo
// 	fmt.Println(jugadoresDeUnEquipo)
// }

// func copySlice() {
// 	sliceOriginal := []int{1, 2, 3, 4}
// 	sliceCopia := make([]int, len(sliceOriginal))
// 	// copy copia elemento por elemento y despues retorna el n° de elem copiados
// 	n := copy(sliceCopia, sliceOriginal)
// 	fmt.Println(n, "numeros copiados:", sliceCopia)
// }

// func arrYsliceComoArgs() {
// 	array1 := [3]int{1, 2, 3}
// 	slice1 := []int{4, 5, 6}
// 	reemplazarPor99(array1, slice1)
// 	fmt.Println("Array:", array1)
// 	fmt.Println("Slice:", slice1)

// }

// func reemplazarPor99(arr [3]int, sli []int) {
// 	arr[0] = 99
// 	sli[0] = 99
// 	// el array al funcionar por valor, se trabaja con una copia del array original
// 	// por lo que, el original no se modifica

// 	// el slice al funcionar por referencia, modifica directamente el slice original

// 	// igualmente la forma correcta de hacer esta funcion seria asi:
// 	// arr[0] = 99
// 	// if len(sli) > 0 {
// 	// 	sli[0] = 00
// 	// }
// 	// como no podes asegurar que un slice tenga una longitud concreta
// 	// (no sabes siquiera si tiene un elemento), es aconsejable comprobarlo primero
// 	// con un if para que el programa no falle durante la ejecucion
// }

// vistas
// func viewsSlice() {
// 	slice1 := []int{1, 10, 3, 4, 5}
// 	view1 := slice1[1:4]
// 	view1[0] = 2
// 	view2 := slice1[0:2]
// 	view3 := slice1[3:]
// 	view4 := slice1[:]

// 	fmt.Println(view1)
// 	fmt.Println(view2)
// 	fmt.Println(view3)
// 	fmt.Println(view4)
// }

// func forOptimizado() {
// 	juegos := []string{"crash", "mario", "winning eleven", "tekken"}
// 	for i, juego := range juegos {
// 		fmt.Println(i, ":", juego)
// 	}

// 	// si no queres indicar el indice se escribe asi
// 	for _, juego := range juegos {
// 		fmt.Println(juego)
// 	}

// }

func funcionSumaConVariosArgs(numeros ...int) int {
	sumaTotal := 0
	for _, numero := range numeros {
		sumaTotal += numero
	}
	return sumaTotal
}
