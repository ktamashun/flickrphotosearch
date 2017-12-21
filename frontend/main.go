package main

import (
	"html/template"
	"net/http"
	"os"
	"fmt"
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
		BackendUrl: os.Getenv("API_ENDPOINT"),
	}

	fmt.Println(fmt.Sprintf("Api endpoint is: %s", os.Getenv("API_ENDPOINT")))

	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, page)
}
