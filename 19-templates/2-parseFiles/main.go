package main

import (
	"fmt"
	"os"
	"text/template"
)

type User struct {
	Name string
	Age  int
}

func main() {
	tmpl, err := template.New("example.txt").ParseFiles("templates/example.txt")
	if err != nil {
		fmt.Println("error creating template")
		panic(err)
	}

	user := User{"lucas", 28}

	tmpl.Execute(os.Stdout, user)
}
