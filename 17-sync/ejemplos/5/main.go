package main

import (
	"fmt"
	"sync"
)

type Book struct {
	name   string
	readed bool
}

func main() {
	var books []Book
	ch := make(chan *Book, 5)
	flagCh := make(chan struct{})
	wg := sync.WaitGroup{}
	num := 10
	for i := 0; i < num; i++ {
		books = append(books, Book{fmt.Sprintf("Book n°%d", i+1), false})
	}

	for i := 0; i < len(books); i++ {
		wg.Add(1)
		go sendBook(&books[i], ch, &wg)
	}
	go readBooks(ch, flagCh) // empieza a leer el canal
	wg.Wait()                // espera a que todas las gorutinas escriban en el ch
	close(ch)                // cierra el canal, termina el range y envia la señal de struct vacio
	<-flagCh                 // espera la señal de struct vacio para seguir con la ejec de main
	// time.Sleep(time.Second * 5)
	fmt.Println(books)
}

func sendBook(book *Book, ch chan<- *Book, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- book
}

func readBooks(ch <-chan *Book, flagCh chan struct{}) {
	for book := range ch {
		book.readed = true
	}
	flagCh <- struct{}{}
}
