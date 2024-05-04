package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	comida := []string{"caramelo", "hamburguesa", "pancho"}
	ch := make(chan string, 10)
	for i := 0; i < len(comida); i++ {
		ch <- comida[i]
	}

	wg.Add(3)
	go comer("jose", ch, &wg)
	go comer("marcos", ch, &wg)
	go comer("antonia", ch, &wg)

	wg.Wait()
}

func comer(persona string, ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(persona, "comio", <-ch)
	// for comida := range ch {
	// 	fmt.Println(persona, "comio", comida)
	// }
}
