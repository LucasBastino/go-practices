package main

import (
	f "fmt"

	"github.com/LucasBastino/practicas-go/sliceMethods"
)

func main() {
	f.Println("slice methods exercises:")
	slice1 := []string{"a", "b", "c", "d"}
	slice2 := []int{1, 2, 3, 4, 5}
	f.Println("slice1 antes de las funciones:", slice1)
	f.Println("PrintSlice:")
	sliceMethods.PrintSlice(slice1)
	f.Println("PrintSliceLine:")
	sliceMethods.PrintSliceLine(slice1)
	// sliceMethods.PrintSliceLineV2(slice1...)
	f.Println("FindIndex:")
	f.Println(sliceMethods.FindIndex(slice1, "a"))
	f.Println("FindIndex and DeleteOneByIndex:")
	slice1 = sliceMethods.DeleteOneByIndex(slice1, sliceMethods.FindIndex(slice1, "a"))
	f.Println(slice1)
	f.Println("Push: a")
	slice1 = sliceMethods.Push(slice1, "a")
	f.Println(slice1)
	// slice1 = sliceMethods.DeleteOneByIndex(slice1, len(slice1)-1)
	slice1, elementoPop1 := sliceMethods.Pop(slice1)
	f.Println(elementoPop1, slice1)
	slice2, elementoPop2 := sliceMethods.Pop(slice2)
	f.Println(elementoPop2, slice2)
	slice1, elementoShift1 := sliceMethods.Shift(slice1)
	f.Println(elementoShift1, slice1)
	slice2, elementoShift2 := sliceMethods.Shift(slice2)
	f.Println(elementoShift2, slice2)
	elemento1 := "H"
	slice1 = sliceMethods.Unshift(slice1, elemento1)
	elemento2 := 17
	slice2 = sliceMethods.Unshift(slice2, elemento2)
	f.Println(slice1, slice2)
}
