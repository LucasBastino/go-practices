package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	ch := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go generarRandom(ch, &wg)
	go generarRandom(ch, &wg)

	go func() {
		for str := range ch {
			fmt.Println(str)
		}
	}()
	wg.Wait()
	close(ch) // hace que el for range deje de iterar
	fmt.Println("funcion finalizada")
}
func generarRandom(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- fmt.Sprintf("numero random: %d", rand.Intn(100))
}
