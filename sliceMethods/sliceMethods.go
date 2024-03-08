package sliceMethods

import "fmt"

func Mostrar2() {
	fmt.Println("muestro 2 desde slicemethdosss")
}

// hacer un delete desde hasta
// y hacer un delete con muchos argumentos
// func delete(slice, elementos... ){

// }

// primero hago uno de a 1
// func deleteOne(slice []int, i int) []int {
// 	return append(slice[:i], slice[i+1:]...)
// }

func deleteOne(s S, i int) {
	return append(slice[:i], slice[i+1:]...)
}
