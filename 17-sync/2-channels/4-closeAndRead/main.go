package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	// wg := sync.WaitGroup{}
// 	ch := make(chan int, 11)
// 	for i := 0; i < 10; i++ {
// 		enviarDato(ch, i)
// 	}
// 	close(ch)
// 	leerDatos(ch)
// }

// func enviarDato(ch chan<- int, i int) {
// 	ch <- i
// 	fmt.Println("enviar dato:", i)
// }

// func leerDatos(ch <-chan int) {
// 	for dato := range ch {
// 		fmt.Println("leer dato:", dato)
// 	}
// }

// lo de abajo es mas rapido porque se esta enviando en paralelo

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int, 11)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go enviarDato(ch, i, &wg)
	}
	wg.Wait()
	close(ch)
	leerDatos(ch)
}

func enviarDato(ch chan<- int, i int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- i
	fmt.Println("enviar dato:", i)
}

func leerDatos(ch <-chan int) {
	for dato := range ch {
		fmt.Println("leer dato:", dato)
	}
}
