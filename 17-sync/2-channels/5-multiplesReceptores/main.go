package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	comida := []string{"caramelo", "hamburguesa", "pancho", "pebete"}
	ch := make(chan string, 10)
	for i := 0; i < len(comida); i++ {
		ch <- comida[i]
		time.Sleep(500 * time.Millisecond)
	}

	wg.Add(4)
	go comer("jose", ch, &wg)
	go comer("marcos", ch, &wg)
	go comer("antonia", ch, &wg)
	go comer("masi", ch, &wg)

	wg.Wait()
}

func comer(persona string, ch <-chan string, wg *sync.WaitGroup) {
	// fmt.Println(persona, "comio", <-ch)
	for comida := range ch {
		fmt.Println(persona, "comio", comida)
		wg.Done()
	}
}
