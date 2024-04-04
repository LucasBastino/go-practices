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

func DeleteOneByValue[typeSlice ~[]c, c comparable](varSlice typeSlice, valueToFind c) typeSlice {
	indexToDelete := FindIndex(varSlice, valueToFind)
	return DeleteOneByIndex(varSlice, indexToDelete)
}

// func DeleteOneByField[typeSlice ~[]e, e any](varSlice typeSlice, field string, valueToFind any) typeSlice {
// 	for i, value := range varSlice {
// 		if reflect.ValueOf(value).FieldByName(field).String() == valueToFind {
// 			varSlice = DeleteOneByIndex(varSlice, i)
// 		}
// 	}
// 	return varSlice
// }

func FindIndexByField[typeSlice ~[]e, e any](varSlice typeSlice, field string, valueToFind any) int {
	tipoDeDato := reflect.TypeOf(valueToFind).String()
	var indexFound int
	switch tipoDeDato {
	case "int":
		fmt.Println("tipo de dato INT")
		for i, value := range varSlice {
			// Convierte un reflect value en int64 y despues en int
			if valueToFind == int(reflect.ValueOf(value).FieldByName(field).Int()) {
				indexFound = i
			} else {
				indexFound = -1
			}
		}
	case "string":
		fmt.Println("ES STRING")
		for i, value := range varSlice {
			// Refleja el valor del campo solicitado de la variable y lo transforma a string
			if valueToFind == reflect.ValueOf(value).FieldByName(field).String() {
				indexFound = i
			} else {
				indexFound = -1
			}
		}
	case "bool":
		fmt.Println("es boolean")
		for i, value := range varSlice {
			if valueToFind == reflect.ValueOf(value).FieldByName(field).Bool() {
				indexFound = i
			} else {
				indexFound = -1
			}

		}
	case "float64":
		fmt.Println("es float")
		for i, value := range varSlice {
			valueToCompare := reflect.ValueOf(value).FieldByName(field).Float()
			fmt.Printf("%T %v\n", valueToCompare, valueToCompare)
			if valueToCompare == valueToFind {
				indexFound = i
			} else {
				indexFound = -1
			}
		}
	default:
		fmt.Println("no es un tipo soportado")
	}
	return indexFound
}

func DeleteOneByField[typeSlice ~[]e, e any](varSlice typeSlice, field string, valueToDelete string) {
	i := FindIndexByField(varSlice, field, valueToDelete)
	fmt.Println(i)
}

// type comparable sirve para comparar cualquier tipo (int, string, etc)
// pero no para operar
// type any sirve para operar con cualquier tipo (int, string, etc)
// pero no para comparar

func FindIndex[typeSlice ~[]c, c comparable](varSlice typeSlice, valueToFind c) int {
	for i, value := range varSlice {
		if value == valueToFind {
			return i
		}
	}
	fmt.Println("Coincidence not found")
	return -1
}

// func FindIndex2[typeSlice ~[]c, c comparable](varSlice typeSlice, value c) int {
// 	for i := 0; i < len(varSlice); i++ {
// 		if varSlice[i] == value {
// 			return i
// 		}
// 	}
// 	fmt.Println("No se encontro coincidencia")
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
