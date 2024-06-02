package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	tmpl, err := template.New("example.txt").Parse("mi nombre es {{.}}")
	if err != nil {
		fmt.Println("error creating template")
		panic(err)
	}

	tmpl.Execute(os.Stdout, "lucas")
}
