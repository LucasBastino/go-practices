// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"sync"
// 	"time"
// )

// type Book struct {
// 	name   string
// 	readed bool
// }

// func main() {
// 	wg := sync.WaitGroup{}
// 	var books []Book
// 	for i := 0; i < 10; i++ {
// 		book := Book{name: fmt.Sprint("Libro ", i+1), readed: false}
// 		books = append(books, book)
// 	}

// 	fmt.Println("Reading books in normal mode:")
// 	for i := 0; i < len(books); i++ {
// 		readBookNormal(&books[i])
// 	}

// 	fmt.Printf("\nReading books in sync mode:\n")
// 	for i := 0; i < len(books); i++ {
// 		wg.Add(1)
// 		go readBookSync(&books[i], &wg)
// 	}
// 	wg.Wait()
// }
// func readBookNormal(book *Book) {
// 	book.readed = true
// 	delay := rand.Intn(1000)
// 	timeSleeped := time.Millisecond * time.Duration(delay)
// 	time.Sleep(timeSleeped)
// 	fmt.Printf("Book readed: %s in %v\n", book.name, timeSleeped)
// }

// func readBookSync(book *Book, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	book.readed = true
// 	delay := rand.Intn(1000)
// 	timeSleeped := time.Millisecond * time.Duration(delay)
// 	time.Sleep(timeSleeped)
// 	fmt.Printf("Book readed: %s in %v\n", book.name, timeSleeped)
// }

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Book struct {
	name   string
	readed bool
}

func main() {
	wg := sync.WaitGroup{}
	var books []Book
	for i := 0; i < 10; i++ {
		books = append(books, Book{fmt.Sprintf("Book %v", i+1), false})
	}

	for i := 0; i < len(books); i++ {
		wg.Add(1)
		go readBook(&books[i], &wg)
	}
	wg.Wait()
	fmt.Println("All books readed")
}

func readBook(book *Book, wg *sync.WaitGroup) {
	defer wg.Done()
	delay := rand.Intn(500)
	timeSleeped := time.Millisecond * time.Duration(delay)
	time.Sleep(timeSleeped)
	book.readed = true
	fmt.Printf("%s readed? %v, in %v\n", book.name, book.readed, timeSleeped)
}
