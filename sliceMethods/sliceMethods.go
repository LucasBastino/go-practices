package sliceMethods

import (
	"errors"
	"fmt"
	"reflect"
)

// data in out
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
	indexToReplace, err := FindIndex(varSlice, oldValue)
	if err != nil {
		fmt.Printf("An error has ocurred in Replace: %v\n", err)
	}
	part1 := append(varSlice[:indexToReplace], newValue)
	part2 := varSlice[indexToReplace+1:]
	return append(part1, part2...)
}

// searcher
func FindIndex[typeSlice ~[]c, c comparable](varSlice typeSlice, valueToFind c) (int, error) {
	for i, value := range varSlice {
		if value == valueToFind {
			return i, nil
		}
	}

	return -1, errors.New("coincidence not found")
}

func FindIndexByField[typeSlice ~[]e, e any](varSlice typeSlice, field string, valueToFind any) (int, error) {
	dataType := reflect.TypeOf(valueToFind).String()
	var indexFound int = -1
	var err error = IsNotIn
	switch dataType {
	case "int":
		fmt.Println("tipo de dato INT")
		for i, value := range varSlice {
			// Convierte un reflect value en int64 y despues en int
			if valueToFind == int(reflect.ValueOf(value).FieldByName(field).Int()) {
				indexFound = i
				err = nil
				break
			}
		}
	case "string":
		fmt.Println("ES STRING")
		for i, value := range varSlice {
			// Refleja el valor del campo solicitado de la variable y lo transforma a string
			if valueToFind == reflect.ValueOf(value).FieldByName(field).String() {
				indexFound = i
				err = nil
				break
			}
		}
	case "bool":
		fmt.Println("es boolean")
		for i, value := range varSlice {
			if valueToFind == reflect.ValueOf(value).FieldByName(field).Bool() {
				indexFound = i
				err = nil
				break
			}
		}
	case "float64":
		fmt.Println("es float")
		for i, value := range varSlice {
			valueToCompare := reflect.ValueOf(value).FieldByName(field).Float()
			fmt.Printf("%T %v\n", valueToCompare, valueToCompare)
			if valueToCompare == valueToFind {
				indexFound = i
				err = nil
				break
			}
		}
	default:
		err = errors.New("Is not a supported type")
	}
	return indexFound, err
}

// delete
func DeleteOneByIndex[typeSlice ~[]s, s any](varSlice typeSlice, index int) typeSlice {
	if index == -1 {
		return varSlice
	} else {
		return append(varSlice[:index], varSlice[index+1:]...)
	}
}

func DeleteOneByValue[typeSlice ~[]c, c comparable](varSlice typeSlice, valueToFind c) typeSlice {
	indexToDelete, err := FindIndex(varSlice, valueToFind)
	if err != nil {
		fmt.Printf("An error has ocurred in DeleteOneByValue: %v\n", err)
	}
	return DeleteOneByIndex(varSlice, indexToDelete)
}

func DeleteOneByField[typeSlice ~[]e, e any](varSlice typeSlice, field string, valueToDelete string) typeSlice {
	i, err := FindIndexByField(varSlice, field, valueToDelete)
	if err != nil {
		fmt.Printf("An error has ocurred in DeleteOneByField: %v\n", err)
	}
	varSlice = DeleteOneByIndex(varSlice, i)
	fmt.Println(i)
	return varSlice
}

type SliceError struct {
	msg string
}

func (se SliceError) Error() string {
	return fmt.Sprintf("An error has ocurred in SliceMethods: %v\n", se.msg)
}

var IsNotIn = SliceError{msg: "The element is not in the slice."}
var Random = SliceError{msg: "Random test error."}

// type comparable sirve para comparar cualquier tipo (int, string, etc)
// pero no para operar
// type any sirve para operar con cualquier tipo (int, string, etc)
// pero no para comparar

// replace, revers, sort (max o min), equal

// hacer un delete desde hasta
// y hacer un delete con muchos argumentos
// func delete(slice, elementos... ){

// }

//  conservar
// conservar desde hasta
