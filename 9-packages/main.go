package main

import (
	//  aca hace falta escribirle test1 porque el package no tiene el mismo nombre que la carpeta
	"fmt"

	test1 "github.com/LucasBastino/practicas-go/folderTestPackage"
	"github.com/LucasBastino/practicas-go/sliceMethods"
)

// y aca no pone nada porque tiene el mismo nombre

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	test1.ImprimirDesdeTest1(slice1)
	sliceCambiado, pop := sliceMethods.Pop(slice1)
	fmt.Println("pop desde sliceMethods", pop, sliceCambiado)
}
