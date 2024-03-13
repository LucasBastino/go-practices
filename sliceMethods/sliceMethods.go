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
	if index == -1 {
		return varSlice
	} else {
		return append(varSlice[:index], varSlice[index+1:]...)
	}

}

// func DeleteOneByValue[typeSlice ~[]s, s any](varSlice typeSlice, valueToDelete any) typeSlice {
// 	var slice []any
// 	slice = append(slice, varSlice...)
// 	var indexToDelete int
// 	for index := range varSlice {
// 		if varSlice[index] == valueToDelete {
// 			indexToDelete = index
// 		}
// 	}
// 	return DeleteOneByIndex(varSlice, indexToDelete)
// }

// type comparable sirve para comparar cualquier tipo (int, string, etc)
// pero no para operar
// type any sirve para operar con cualquier tipo (int, string, etc)
// pero no para comparar
func FindIndex[typeSlice ~[]c, c comparable](varSlice typeSlice, value c) int {
	for i := 0; i < len(varSlice); i++ {
		if varSlice[i] == value {
			return i
		}
	}
	fmt.Println("No se encontro coincidencia")
	return -1
}

// func FindIndex2[typeSlice ~[]E, E comparable](varSlice typeSlice, value E) int {
// 	for i, valor := range varSlice {
// 		if valor == value {
// 			return i
// 		}
// 	}
// 	return -1
// }

//  es innecesario, funciona igual que append
func Push[typeSlice ~[]E, E any](varSlice typeSlice, elemento E) typeSlice {
	return append(varSlice, elemento)
}

func Pop[typeSlice ~[]S, S any](varSlice typeSlice) (typeSlice, any) {
	elemento := varSlice[len(varSlice)-1]
	varSlice = DeleteOneByIndex(varSlice, len(varSlice)-1)
	return varSlice, elemento
}

func Shift[typeSlice ~[]S, S any](varSlice typeSlice) (typeSlice, any) {
	elemento := varSlice[0]
	varSlice = DeleteOneByIndex(varSlice, 0)
	return varSlice, elemento
}

func Unshift[typeSlice ~[]E, E any](varSlice typeSlice, elemento E) typeSlice {
	varSlice2 := []E{elemento}
	return append(varSlice2, varSlice...)
}

// hacer un delete desde hasta
// y hacer un delete con muchos argumentos
// func delete(slice, elementos... ){

// }

//  conservar
// conservar desde hasta
