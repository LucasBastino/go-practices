package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int, 2)

	for i := 0; i < 10; i++ {
		wg.Add(2)
		go enviarDato(ch, i, &wg)
		go recibirDato(ch, &wg)
	}
	wg.Wait()
}

func enviarDato(ch chan<- int, i int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- i
	fmt.Println("dato enviado:", i)
}

func recibirDato(ch <-chan int, wg *sync.WaitGroup) {
	time.Sleep(time.Duration(1000) * time.Millisecond)
	defer wg.Done()
	fmt.Println("dato recibido:", <-ch)
}

// aca aparece algunas veces que se recibe primero el dato y despues se lo envia
// eso es solamente por que el Println de enviar se hace DESPUES de enviar el dato
// entonces algunas veces se recibe el dato mas rapido de lo que se IMPRIME EL ENVIO
