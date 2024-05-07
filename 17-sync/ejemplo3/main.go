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
	// close(ch)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for str := range ch {
			fmt.Println(str)
		}
	}()
	wg.Wait()
}
func generarRandom(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- fmt.Sprintf("numero random: %d", rand.Intn(100))
}
