package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	numero := 10
	err1 := fmt.Sprintf("error numero %v con Sprintf", numero)
	err2, _ := fmt.Println("hola", numero)
	err3 := fmt.Errorf("error numero %v con Errorf", numero)
	err4 := errors.New("error con Errors.New()")
	var err5 error
	fmt.Printf("error con Sprintf es de tipo %T\n", err1)
	fmt.Printf("error con Println es de tipo %T\n", err2)
	fmt.Printf("error con Errorf es de tipo %T\n", err3)
	fmt.Printf("error con errors.New() es de tipo %T\n", err4)
	fmt.Printf("error con error es de tipo %T\n", err5)
	mostrandoError()

}

// los unicos que retornar un error son errors.New() y fmt.Errorf()

// generalmente se usa asi
func mostrandoError() {
	numero := 5
	log.Println(fmt.Errorf("imprimo el error numero"))
	log.Println(errors.New(fmt.Sprintf("imprimo el error numero %v", numero)))

	// panic(fmt.Errorf("imprimo el error numero %v", numero))
	// panic(errors.New(fmt.Sprintf("imprimo el error numero %v", numero)))

	// para mi fmt.Errorf es cuando hay que usar variables para imprimir
	// y errors.New es para cuando tenes que imprimir un string comun
}
