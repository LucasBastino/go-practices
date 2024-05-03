package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {

		ch <- "Hola"
	}()

	strRecibido := <-ch
	fmt.Println(strRecibido)
}

//  se ejecuta tanto el main como la gorrutina a la vez, por lo que main se queda bloqueada escuchando al canal
//  hasta recibir el dato a guardar en strRecibido
