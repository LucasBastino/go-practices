package main

import "fmt"

func main() {
	fmt.Println("Ejercicios de variables:")
	definoVariables()
	// muestroTiposYValor()
	// leerYMostrarDatos()
	// leerYMostrarDatosAvanzado()
	// leerYmostrarDatos2()
	// concatenar()
	// mostrarShadowing()
}

func definoVariables() {
	var nombre string = "Jose"
	apellido := "Ramirez"
	var nota1 int = 5
	nota2 := 8
	promedio := (float32(nota1+nota2) / 2)

	fmt.Println("\nDefino variables:", nombre, apellido, nota1, nota2, promedio)
}

// func muestroTiposYValor() {
// 	nombre := "Lucila"
// 	apellido := "Argarañaz"
// 	nota1 := 10
// 	nota2 := 7
// 	promedio := (float32(nota1+nota2) / 2)

// 	fmt.Printf(`Muestro tipo y valor:
// 	%v %T
// 	%v %T
// 	%v %T
// 	%v %T
// 	%v %T

// 	`,
// 		nombre, nombre, apellido, apellido, nota1, nota1, nota2, nota2, promedio, promedio)
// }

// func leerYMostrarDatos() {

// 	var nombre string
// 	var edad int

// 	fmt.Println("Como te llamas?")
// 	fmt.Scan(&nombre)
// 	fmt.Println("Cuantos años tenes?")
// 	fmt.Scan(&edad)

// 	fmt.Println("El alumno se llama", nombre, "y tiene", edad, "años")
// }

// func leerYMostrarDatosAvanzado() {

// 	var hora int
// 	var minuto int
// 	var segundo int

// 	fmt.Println("Que hora es? (en formato HH:MM:SS)")
// 	fmt.Scanf("%d:%d:%d", &hora, &minuto, &segundo)

// 	fmt.Printf("Hora: %v Minuto: %v Segundo: %v\n", hora, minuto, segundo)
// }

// func leerYmostrarDatos2() {
// 	var nombre string
// 	var edad int
// 	var estatura float32

// 	fmt.Println("Ingresa nombre, edad y estatura:")
// 	fmt.Scanf("%s %d %f", &nombre, &edad, &estatura)
// 	fmt.Printf("%v tiene %v años y mide %vcm\n", nombre, edad, estatura)
// }

// func concatenar() {
// 	frase1 := "hola"
// 	frase2 := "chau"
// 	fraseConcatenada := frase1 + " " + frase2
// 	fmt.Println(fraseConcatenada)
// }

// shadowing u ocultacion (scope para mi)
// func mostrarShadowing() {
// 	a := 0
// 	b := 0

// 	if true {
// 		a := 1 // aca estas declarando una nueva variable, aunque tenga el mismo nombre es una nueva y solo existe en el if
// 		b = 1  // aca estas asignandole el valor 0 a la variable b ya declarada afuera del if, por lo tanto, la afecta
// 		a++
// 		b++
// 		fmt.Printf("adentro del if a = %v, b = %v\n", a, b) // esto va a imprimir a = 2 y b = 2
// 	}

// 	fmt.Printf("afuera del if a = %v, b = %v\n", a, b) // esto va a imprimir a = 0 y b = 2
// }
