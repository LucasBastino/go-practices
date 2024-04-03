package sliceMethods

import (
	"fmt"
	"reflect"
)

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

func DeleteOneByValue[typeSlice ~[]c, c comparable](varSlice typeSlice, valueToDelete c) typeSlice {
	indexToDelete := FindIndex(varSlice, valueToDelete)
	return DeleteOneByIndex(varSlice, indexToDelete)
}

func DeleteOneByField[typeSlice ~[]e, e any](varSlice typeSlice, field string, valueToDelete any) typeSlice {
	for _, value := range varSlice {
		valor1 := reflect.ValueOf(value).FieldByName(field).Interface().(string)
		fmt.Println(valor1, valueToDelete)
		if string(valor1) == valueToDelete {
			fmt.Println("son iguales")
		}
		// // fmt.Println(i, reflect.ValueOf(value).FieldByName(field))
		// if reflect.ValueOf(value).FieldByName(field) == valueToDelete {
		// 	// varSlice = DeleteOneByIndex(varSlice, i)
		// 	fmt.Println("el valor buscado", valueToDelete, "es igual a el valor del campo:", field)
		// }
	}
	return varSlice
}

func ShowField[typeSlice ~[]e, e any](varSlice typeSlice, field string) {
	for i, value := range varSlice {
		fmt.Println(i, reflect.ValueOf(value).FieldByName(field))
	}
}

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

// es innecesario, funciona igual que append
func Push[typeSlice ~[]E, E comparable](varSlice typeSlice, elemento E) typeSlice {
	return append(varSlice, elemento)
}

func Pop[typeSlice ~[]S, S comparable](varSlice typeSlice) (typeSlice, S) {
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

func Replace[typeSlice ~[]E, E comparable](varSlice typeSlice, oldValue E, newValue E) typeSlice {
	indexToReplace := FindIndex(varSlice, oldValue)
	part1 := append(varSlice[:indexToReplace], newValue)
	part2 := varSlice[indexToReplace+1:]
	return append(part1, part2...)

}

// replace, revers, sort (max o min), equal

// hacer un delete desde hasta
// y hacer un delete con muchos argumentos
// func delete(slice, elementos... ){

// }

//  conservar
// conservar desde hasta
