package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan string)
	wg := sync.WaitGroup{}
	flagCh := make(chan struct{})
	num := 5

	for i := 0; i < num; i++ {
		wg.Add(1)
		str := fmt.Sprintf("String numero: %v", i+1)
		go sendString(str, &wg, ch)

	}

	go readStrings(flagCh, ch)
	wg.Wait()
	close(ch)
	<-flagCh

	fmt.Println("funcion finalizada")

}

func sendString(str string, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	ch <- str
}

func readStrings(flagCh chan struct{}, ch chan string) {
	for str := range ch {
		fmt.Println(str)
	}
	flagCh <- struct{}{}
}
