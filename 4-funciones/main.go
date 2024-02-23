package main

import "fmt"

func main() {
	fmt.Println("Ejercicios de funciones:")
	// fmt.Println(funcionNormal(5, 3))
	// fmt.Println(funcionMultiplesValores(3, 8))
	// fmt.Println(funcionMultiplesValoresNombrados(10, 5))
	// asignandoValoresAnteriorFuncion()
	// descartandoValores()
	// ejDemostrativoFuncionPorValor()
	// ejDemostrativoFuncionPorReferencia()
	// funcionProbarScopes()
	// funcionesLiterales()
	// probarFuncionInmediata()
	// pasajeArgumentosDistintosTipos()
}

// func funcionComun(a int, b int) int { // (a, b int) es la version simplificada
// 	// devuelve el numero maximo
// 	if a > b {
// 		return a
// 	} else {
// 		return b
// 	}
// }

// func funcionMultiplesValores(a, b int) (int, int) {
// 	// devuelve el maximo y despues el minimo
// 	if a > b {
// 		return a, b
// 	} else {
// 		return b, a
// 	}
// }

// func funcionMultiplesValoresNombrados(a, b int) (max int, min int) {
// 	if a > b {
// 		max = a
// 		min = b
// 	} else {
// 		max = b
// 		min = a
// 	}
// 	return
// }

// func asignandoValoresAnteriorFuncion() {
// 	numeroMaximo, numeroMinimo := funcionMultiplesValoresNombrados(11, 5)
// 	fmt.Printf("El numero maximo es %v y el minimo es %v\n", numeroMaximo, numeroMinimo)
// }

// func descartandoValores() {
// 	// el "_" descarta el valor retornado y no lo asigna, en este caso, descarta el segundo valor
// 	//  esto es necesario ya que Go no permite variables que no se van a usar
// 	numeroMaximo, _ := funcionMultiplesValoresNombrados(5, 4)
// 	fmt.Println("El numero mas grande es", numeroMaximo)
// }

// func incrementaPorValor(a int) {
// 	a++
// }

// func ejDemostrativoFuncionPorValor() {
// 	a := 2
// 	fmt.Println(a)
// 	incrementaPorValor(a)
// 	fmt.Println(a)
// 	// esto va a imprimir:
// 	// a = 2
// 	// a = 2
// 	// ya que en incrementaValor(a int) estas trabajando con una copia de "a"
// }

// func incrementaPorReferencia(a *int) {
// 	*a++
// }

// func ejDemostrativoFuncionPorReferencia() {
// 	a := 2
// 	fmt.Println(a)
// 	fmt.Println(&a)
// 	incrementaPorReferencia(&a)
// 	fmt.Println(a)
// 	// esto va a imprimir:
// 	// a = 2
// 	// a = 3
// 	// ya que incrementa el valor de la variable a la
// 	// cual "a" esta apuntando
// }

// func funcionProbarScopes() {
// 	a := 1
// 	fmt.Println(a)
// 	if true {
// 		a := 0
// 		fmt.Println(a)
// 		// imprime 0 porque "a" esta declarada adentro del if,
// 		//  es una variable aparte
// 	}
// 	if true {
// 		a++
// 	}
// 	fmt.Println(a)
// 	// imprime 2 porque el anterior if incrementa
// 	//  el "a" del principio de la funcion, no declara una nueva
// }

// func funcionesLiterales() {
// 	var operador func(int, int) int // declarar como una variable normal, sin el =
// 	operador = sumar
// 	fmt.Println(operador(4, 2))
// 	operador = multiplicar
// 	fmt.Println(operador(4, 2))
// 	operador = restar
// 	fmt.Println(operador(4, 2))
// 	operador = dividir
// 	fmt.Println(operador(4, 2))
// }

// func sumar(a, b int) int {
// 	return a + b
// }

// func multiplicar(a, b int) int {
// 	return a * b
// }

// func restar(a, b int) int {
// 	return a - b
// }

// func dividir(a, b int) int {
// 	return a / b
// }

// // funcion que se crea y se invoca justo despues, sin necesidad de asignarsela a una variable
// func() {
// 	fmt.Print("funcion que se crea y se invoca")
// }()
// fijarse que tiene () al final para que se ejecute
// solo puede ejecutarse adentro de otra funcion o adentro de main

// func probarFuncionInmediata() {
// 	func() {
// 		fmt.Print("funcion que se crea y se invoca")
// 	}()
// }

// func pasajeArgumentosDistintosTipos() {

// 	// imprimirInt32(int32("8"))
// 	// esto no se puede hacer asi, para convertir un string en un int32
// 	// hay que usar la funcion strconv.Atoi(variable)

// 	var numero int64 = 8
// 	imprimirInt32(int32(numero))
// 	// antes de pasar el argumento hay que convertirlo al tipo
// 	// requerido en esa funcion, si mandas un int64 no compila
// }

// func imprimirInt32(numero int32) {
// 	fmt.Println(numero)
// }
