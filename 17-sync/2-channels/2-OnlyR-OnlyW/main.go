package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go escribirDato(i, ch, &wg)
		go leerDatos(ch, &wg)
	}
	wg.Wait()
}

func leerDatos(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("dato recibido:", <-ch)

}

func escribirDato(i int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("dato enviado:", i)
	ch <- i
}

// en los canales sin buffer cuando se envia un dato la gorutina queda bloqueada hasta que
// alguien reciba el dato del canal

// func main() {
// 	// wg := sync.WaitGroup{}
// 	ch := make(chan int)

// 	go enviarDato(ch)
// 	recibirDatos(ch)
// }

// func enviarDato(ch chan<- int) {
// 	for i := 0; i < 10; i++ {

// 		ch <- i
// 		fmt.Println("dato enviado:", i)
// 	}
// }

// func recibirDatos(ch <-chan int) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println("dato recibido:", <-ch)
// 	}
// }
