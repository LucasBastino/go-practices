package main

import (
	f "fmt"

	"github.com/LucasBastino/practicas-go/sliceMethods"
)

func main() {
	f.Println("slice methods exercises:")
	slice1 := []string{"a", "b", "c", "d"}
	f.Println("slice1 antes de las funciones:", slice1)
	f.Println("PrintSlice:")
	sliceMethods.PrintSlice(slice1)
	f.Println("PrintSliceLine:")
	sliceMethods.PrintSliceLine(slice1)
	// sliceMethods.PrintSliceLineV2(slice1...)
	f.Println("DeleteOneByIndex")
	slice1 = sliceMethods.DeleteOneByIndex(slice1, 0)
	f.Println(slice1)
	slice1 = sliceMethods.DeleteOneByValue(slice1, "d")
	f.Println("DeleteOneByValue: 'c'")
	f.Println(slice1)
}
