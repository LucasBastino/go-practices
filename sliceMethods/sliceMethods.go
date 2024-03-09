package sliceMethods

import "fmt"

// metodo especifico para testear en 9-packages
func ImprimirDesdeSliceMethods[typeSlice ~[]s, s any](varSlice typeSlice) {
	fmt.Println(varSlice, "desde sliceMethods")
}

// metodos generales
func PrintSlice[Slice ~[]sli, sli any](slice Slice) {
	fmt.Println(slice)
}

func PrintSliceLine[typeSlice ~[]s, s any](varSlice typeSlice) {
	for _, element := range varSlice {
		fmt.Println(element)
	}
}

// func PrintSliceLineV2(letras ...string) {
// 	for _, letra := range letras {
// 		fmt.Println(letra)
// 	}
// }

func DeleteOneByIndex[typeSlice ~[]s, s any](varSlice typeSlice, index int) typeSlice {
	return append(varSlice[:index], varSlice[index+1:]...)
}

func DeleteOneByValue[typeSlice ~[]s, s any](varSlice typeSlice, valueToDelete any) typeSlice {
	var indexToDelete int
	for index := range varSlice {
		if varSlice[index] == valueToDelete {
			indexToDelete = index
		}
	}
	return DeleteOneByIndex(varSlice, indexToDelete)
}

// hacer un delete desde hasta
// y hacer un delete con muchos argumentos
// func delete(slice, elementos... ){

// }

//  conservar
// conservar desde hasta
