package main

import (
	"net/http"
	"text/template"
)

type User struct {
	Name string
	Age  int
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))

}

func main() {
	muxer := http.NewServeMux()

	svr := &http.Server{
		Addr:    ":8085",
		Handler: muxer,
	}

	muxer.HandleFunc("/index", index)
	muxer.HandleFunc("/about", about)
	svr.ListenAndServe()

}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func about(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "about.html", nil)
}
