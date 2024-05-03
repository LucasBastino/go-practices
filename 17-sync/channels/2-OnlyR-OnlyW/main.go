package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(11)
	for i := 0; i < 10; i++ {
		go escribirDato(i, ch, &wg)
	}
	go leerDatos(ch, &wg)
	wg.Wait()
}

func leerDatos(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Println(num)
	}

}

func escribirDato(i int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- i
}
