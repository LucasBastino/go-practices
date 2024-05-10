package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// func main() {
// 	intCh := make(chan int)
// 	flagCh := make(chan struct{})
// 	// wg := sync.WaitGroup{}
// 	// wg.Add(2)
// 	go generarMuchosRandom(3, intCh, flagCh)

// 	go func() {
// 		for num := range intCh {
// 			time.Sleep(time.Second)
// 			fmt.Println("el numero recibido es ", num)
// 		}

// 	}()
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(i, "segundo")
// 		time.Sleep(time.Second)
// 	}
// 	<-flagCh // es como un close(ch)? no se
// 	fmt.Println("funcion finalizada")
// }

// func generarMuchosRandom(num int, ch chan int, flagCh chan struct{}) {
// 	for i := 0; i < num; i++ {
// 		generarRandom(ch)
// 	}
// 	fmt.Println("generarMuchos termino")
// 	flagCh <- struct{}{}
// }
// func generarRandom(ch chan int) {
// 	ch <- rand.Intn(100)
// }

func main() {
	intCh := make(chan int)
	wg := sync.WaitGroup{}
	num := 5
	wg.Add(num)
	go generarMuchosRandom(num, intCh, &wg)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range intCh {
			time.Sleep(time.Second * 2)
			fmt.Println("el numero recibido es ", num)
		}
		// se puede hacer wg.Add y Done porque en otro lugar del codigo cierro el ch
		// sino, es deadlock, hay que cerrarlo si o si
		fmt.Println("imprimir range termino")
		// esto se imprime si tengo el time.Second * 2
		// pero si esta timeSecond * 1 no se imprime
		// si pongo el for de abajo tampoco se imprime, si se lo saco si
	}()
	for i := 0; i < 15; i++ { // no se para que puse esto pero igual le da tiempo al imprimir range termino
		fmt.Println(i, "segundo")
		time.Sleep(time.Second)
	}
	wg.Wait() // esto no finaliza este ultimo for de imprimir segundos,
	//  solo espera a que terminen todos los wg.Done para seguir con lo de abajo
	// time.Sleep(time.Second * 2) // asi le doy tiempo a que el for imprima el ultimo numero
	fmt.Println("imprimo desde main")

	fmt.Println("funcion finalizada")
}

func generarMuchosRandom(num int, ch chan int, wg *sync.WaitGroup) {
	for i := 0; i < num; i++ {
		generarRandom(ch, wg)
	}
	fmt.Println("generarMuchos termino")
	close(ch) // cierra el canal y por lo tanto, termina el for range

}
func generarRandom(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- rand.Intn(100)
}

// conclusiones:
// el close(ch) este donde esté termina con el for range al cerrar el canal
// (el for igual lee todos los valores y despues cierra)
// y permite que la funcion siga su curso
// se le puede agregar un Wg.Add y Wg.Done a una gorrutina con for range
// siempre y cuando se cierre el canal en algun lugar del codigo (no hace falta que sea en esa funcion)
// el wg.Wait espera a todos los wg.Done y sigue la ejecucion de abajo
// haciendo esto te quedas tranquilo de que se imprime todo lo que se tenga que imprimir
// y no hace falta todo el quilombo con los time.Second que habia hecho arriba
