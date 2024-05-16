package main

import "fmt"

type Book struct {
	name   string
	readed bool
}

func main() {
	var books []Book
	ch := make(chan *Book, 10)
	for i := 0; i < 10; i++ {
		go func() {

			ch <- &Book{fmt.Sprint("Book ", i), false}
		}()
	}

	esperando := putBooks(books, ch)
	<-esperando
	fmt.Println(books)
}

func putBooks(books []Book, ch chan *Book) <-chan struct{} {
	chEspera := make(chan struct{})
	go func() {
		for book := range ch {
			books = append(books, *book)
		}
		fmt.Println("ciclo terminado")
		close(ch) // esto me parece que esta mal, no se deberia cerrar el canal aca
	}()
	return chEspera
}
