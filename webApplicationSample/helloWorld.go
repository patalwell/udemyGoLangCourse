package main

import (
	"io/ioutil"
	"net/http"
	"html/template"
)

type Page struct {
	Title string
	Body []byte
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func helloWorld(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	t, _ := template.ParseFiles("edit.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)
}
