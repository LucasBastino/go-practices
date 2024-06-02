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
	tmpl, err := template.New("newTemplate").ParseFiles("templates/example.txt")
	if err != nil {
		fmt.Println("error creating template")
		panic(err)
	}

	user := User{"marse", 52}
	tmpl.ExecuteTemplate(os.Stdout, "example.txt", user)
}
