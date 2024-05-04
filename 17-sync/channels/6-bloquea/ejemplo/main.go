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
	wg := sync.WaitGroup{}
	ch := make(chan *Book, 10)
	var books []Book
	for i := 0; i < 10; i++ {
		books = append(books, Book{fmt.Sprint("Book ", i+1), false})
	}

	for i := 0; i < len(books); i++ {
		wg.Add(1)
		go sendBook(&books[i], ch, &wg)
	}

	go readBook(ch, &wg)
	wg.Wait()
	fmt.Println(books)

}

func sendBook(book *Book, ch chan<- *Book, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(book.name, "sended")
	ch <- book
}

func readBook(ch <-chan *Book, wg *sync.WaitGroup) {
	defer wg.Done()
	// book := <-ch
	// book.readed = true
	for book := range ch {
		book.readed = true
		fmt.Println(book.name, "readed")
	}

}
