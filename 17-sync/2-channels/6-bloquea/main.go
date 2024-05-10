package main

import (
	"fmt"
	"time"
)

func main() {
	bloquea := funcionParalela()
	fmt.Println("esperando ch con struct{}")
	time.Sleep(700 * time.Millisecond)
	fmt.Println("esperando ch con struct{}")
	<-bloquea
	// ↑ ya es un chan de struct{} desde la primera linea
	//  ahora esta esperando que alguien cierre el ch,
	// cuando se cierre, main ya puede seguir con su ejecucion normalmente
	// se usa como si fuera un Wait del wg
	fmt.Println("finalizado")

}

func funcionParalela() <-chan struct{} {
	ch := make(chan struct{})
	go func() { // funcion cualquiera
		for i := 0; i < 3; i++ {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("numero", i)
		}
		fmt.Println("ciclo for terminado")
		close(ch)
	}()
	return ch
}
