package main

import (
	"html/template"
	"net/http"
)

type Page struct {
	Title,Content string
}

func displayPage(rw http.ResponseWriter,r *http.Request) {
	p := &Page{
		Title: "The Barbeque",
		Content: "Tasty food made here !!",
	}

	t := template.Must(template.ParseFiles("simple.html"))
	t.Execute(rw,p)
}

func main() {
	http.HandleFunc("/food",displayPage)
	http.ListenAndServe(":4000",nil)
}
