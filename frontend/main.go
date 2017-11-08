package main

import (
	"html/template"
	"net/http"
)

type IndexPage struct {
	BackendUrl string
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":80", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	page := &IndexPage{
		BackendUrl: "http://localhost:8101",
	}

	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, page)
}
