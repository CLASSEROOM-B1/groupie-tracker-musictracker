package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to our music tracker website !")
	})

	fmt.Println("Server is running on localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
func OpenFile(w http.ResponseWriter, r *http.Request, filename string) {
	templates, err := template.ParseFiles("../static/" + filename + ".html")
	if err != nil {
		log.Fatal(err)
	}
	templates.Execute(w, templates)
}
func Hostler(w http.ResponseWriter, r *http.Request) {
}
