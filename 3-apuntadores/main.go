package main

import "fmt"

func main() {
	fmt.Println("Ejercicios de apuntadores:")
	// apuntadores()
	// apuntadorANil()
	// comparacionTeoriaApuntadores()
}

// func apuntadores() {
// 	// var p * int, podes hacerlo asi tambien y despues asignarle una direccion
// 	a := 5
// 	p := &a // a la variable "p" se le asigna la direccion de "a", o sea "p" es un apuntador a "a"

// 	fmt.Println("valor de a:", a)                           // muestra el valor de a
// 	fmt.Println("valor de p, que es la direccion de a:", p) // muestra la direccion de memoria de a
// 	fmt.Println("valor de *p:", *p)                         // muestra el valor de la variable a la que apunta "p"

// 	*p = 10
// 	// le asigna el valor 10 a la variable que se encuentra en la direccion a la que apunta "p", en este caso es "a"
// 	// por lo tanto, "a" pasa de valer 5 a valer 10
// 	fmt.Println("ejecuto *p = 10")

// 	fmt.Println("valor de a:", a)                           // muestra el valor de a
// 	fmt.Println("valor de p, que es la direccion de a:", p) // muestra la direccion de memoria de a
// 	fmt.Println("valor de *p:", *p)                         // muestra el valor de la variable a la que apunta "p"
// 	// basicamente *p = a
// }

// func apuntadorANil() {
// 	// si un apuntador apunta a "nil" significa que no apunta a ninguna parte
// 	var p *int
// 	p = nil
// 	// a := 5
// 	// p = &a
// 	if p == nil {
// 		fmt.Println("'p' no apunta a ningun lado")
// 	} else {
// 		fmt.Println("p tiene la direccion", p)
// 	}
// 	// el valor "nil" puede ser utilizado para reiniciar un puntero
// 	// o para comprobar si el puntero apunta a algun lugar valido
// }

// func comparacionTeoriaApuntadores() {
// 	a := 0
// 	b := 0
// 	pa := &a
// 	pb := &b
// 	// pb = &a
// 	// *pb = 5

// 	if pa == pb {
// 		fmt.Println("pa y pb apuntan a la misma direccion")
// 	} else {
// 		fmt.Println("pa y pb apuntan a direcciones diferentes")
// 	}

// 	if *pa == *pb {
// 		fmt.Println("pa y pb apuntan a variables con valores iguales")
// 	} else {
// 		fmt.Println("pa y pb apuntan a variables con valores diferentes")
// 	}
// }
