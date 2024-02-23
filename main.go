package main

func main() {
	// definoVariables()
	// muestroTiposYValor()
	// leerYMostrarDatos()
	// leerYMostrarDatosAvanzado()
	// leerYmostrarDatos2()
	// concatenar()
	// ifYElseIf(25)
	// ingresoYComparo()
	// opciones()
	// switchConCondicional()
	// forContinueYBreak()
	// forNoIterativo()
	// forIterativo()
	// forIterativo2()
	// mostrarShadowing()
	// apuntadores()
	// apuntadorANil()
	// comparacionTeoriaApuntadores()
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
}

// func definoVariables() {
// 	var nombre string = "Jose"
// 	apellido := "Ramirez"
// 	var nota1 int = 5
// 	nota2 := 8
// 	promedio := (float32(nota1+nota2) / 2)

// 	fmt.Println("\nDefino variables:", nombre, apellido, nota1, nota2, promedio)
// }

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

// func ifYElseIf(numero int) {
// 	if numero > 25 {
// 		fmt.Println("el numero es mayor a 25")
// 	} else if numero < 25 {
// 		fmt.Println("el numero es menor a 25")
// 	} else {
// 		fmt.Println("el numero es 25")
// 	}
// }

// func ingresoYComparo() {
// 	fmt.Println("ingrese un numero")
// 	var numero int
// 	fmt.Scan(&numero)
// 	if numero > 25 {
// 		fmt.Println("el numero es mayor a 25")
// 	} else if numero < 25 {
// 		fmt.Println("el numero es menor a 25")
// 	} else {
// 		fmt.Println("el numero es 25")
// 	}
// }

// func opciones() {
// 	var opcion int
// 	fmt.Println("elija una opcion del 1 al 5")
// 	fmt.Scan(&opcion)

// 	switch opcion {
// 	case 1:
// 		fmt.Println("elegiste la opcion 1 campeon")
// 	case 2:
// 		fmt.Println("elegiste la 2 maquina")
// 	case 3:
// 		fmt.Println("has elegido la opcion 3")
// 	case 4:
// 		fmt.Println("te puse")
// 	case 5:
// 		fmt.Println("elegite la 5 pa")
// 	default:
// 		fmt.Println("sos medio boludo")
// 	}
// }

// func switchConCondicional() {
// 	var numero int
// 	fmt.Println("ingresa un numero")
// 	fmt.Scan(&numero)
// 	switch {
// 	case numero < 0:
// 		fmt.Println("el numero es negativo")
// 	case numero == 0:
// 		fmt.Println("el numero es 0")
// 	case numero > 0 && numero < 10:
// 		fmt.Println("el numero es positivo y menor a 10")
// 	case numero >= 10:
// 		fmt.Println("el numero es mayor o igual a 10")
// 	}
// }

// func forContinueYBreak() {
// 	for {
// 		var opcion int
// 		fmt.Println("ingresa 1 para seguir o 2 para salir")
// 		fmt.Scan(&opcion)
// 		if opcion == 1 {
// 			fmt.Println("elegiste seguir")
// 			continue
// 		} else if opcion == 2 {
// 			fmt.Println("elegiste salir")
// 			break
// 		} else {
// 			fmt.Println("elegiste mal la opcion")
// 		}
// 	}
// }

// func forNoIterativo() {
// 	for opcion := "inicio"; opcion != "salir"; fmt.Scan(&opcion) {
// 		fmt.Println(`escribi "salir" para salir, sino sigue`)
// 	}
// }

// func forIterativo() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(i + 1)
// 	}
// }

// func forIterativo2() {
// 	for i := 0; i < 10; i++ {
// 		for j := 0; j < 10; j++ {
// 			fmt.Printf("i: %v, j: %v\n", i, j)
// 		}
// 	}
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

// apuntadores
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
