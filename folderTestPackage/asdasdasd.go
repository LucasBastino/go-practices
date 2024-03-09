package test1

import "fmt"

func ImprimirDesdeTest1[typeSlice ~[]s, s any](varSlice typeSlice) {

	fmt.Println(varSlice, "desde Test1")
}
