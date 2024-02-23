package main

import "fmt"

func main() {
	fmt.Println("Ejercicios de condicionales:")
	// ifYElseIf(25)
	// ingresoYComparo()
	// opciones()
	// switchConCondicional()
	// forContinueYBreak()
	// forNoIterativo()
	// forIterativo()
	// forIterativo2()
}

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
