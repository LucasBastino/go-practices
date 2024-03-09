package sliceMethods

import "fmt"

func ImprimirSlice[Slice ~[]sli, sli any](slice Slice) {
	fmt.Println(slice)
}

// func MostrarElemento[Slice ~[]sli1, sli1 any](slice Slice, elemento int) {
// 	fmt.Println(slice[elemento])
// }

func BorrarElemento[Slice ~[]sli1, sli1 any](slice Slice, elemento int) Slice {
	return append(slice[:elemento], slice[elemento+1:]...)
}

// hacer un delete desde hasta
// y hacer un delete con muchos argumentos
// func delete(slice, elementos... ){

// }

// primero hago uno de a 1
// func deleteOne(slice []int, i int) []int {
// 	return append(slice[:i], slice[i+1:]...)
// }
