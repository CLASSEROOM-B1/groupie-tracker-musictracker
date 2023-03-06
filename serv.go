package main

import (
	"fmt"
	"html/template"
	"log"
	"musicTracker/musicTracker"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		parsedTemplate, _ := template.ParseFiles("./static/templates/map.html")
		Artists := musicTracker.GetAllArtists()
		err := parsedTemplate.Execute(w, Artists)
		if err != nil {
			log.Println("Error executing template :", err)
			return
		}
	})
	fmt.Println("Server is running on localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
