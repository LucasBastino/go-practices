package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	chans := []chan int{ch1, ch2, ch3}

	// pasando por todos los canales
	for i := range chans {
		go func(i int, ch chan int) {
			// inicio un ciclo for infinito por cada canal
			// en el que imprimo el numero de canal
			for {
				time.Sleep(time.Duration(i+1) * time.Second)
				ch <- i + 1
			}
		}(i, chans[i])
	}

	// recibo el dato 10 veces del canal que este listo en el momento
	// si el canal que esta listo es el 1 hago tal cosa, en este caso: recibir un 1
	// si el canal es el 2 hago otra cosa, en este caso, recibo un 2
	for i := 0; i < 10; i++ {
		select {
		case num := <-ch1:
			fmt.Println(num)
		case num := <-ch2:
			fmt.Println(num)
		case num := <-ch3:
			fmt.Println(num)
		}

	}
}
