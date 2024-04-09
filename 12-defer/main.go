package main

func main() {
	// regla1()
	// regla2()
	// regla3()
}

// func regla1() {
// 	i := 11
// 	defer fmt.Println(i)
// 	i = 18
// }

// Defer toma la variable en el momento en que se lo llama, por mas que despues i sea 18, defer imprime 5,
// que es lo que valia i cuando se llamo a defer

// func regla2() {
// 	for i := 0; i < 5; i++ {
// 		defer fmt.Println(i)
// 	}
// }

// Defer toma como regla: ultimo en entrar, primero en salir, por lo tanto aca imprime: 4 3 2 1 0

// func regla3() (i int) {
// 	defer func() { i++ }()
// 	return 1
// }

// como la funcion retorna la variable i y esa funcion tiene un return 1, a i se le asigna el valor 1
// despues hace el defer y el valor de i se incrementa en 1, despues retorna i, que es lo que pide la funcion
//  e i es igual a 2
