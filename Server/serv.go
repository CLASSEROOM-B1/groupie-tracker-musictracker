package musicTracker

import (
	"fmt"
	"html/template"
	"log"
	musicTracker "musicTracker/getAPI"
	"net/http"
)

func Serv() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		parsedTemplate, _ := template.ParseFiles("./static/templates/MainPage.html")
		Artists := musicTracker.GetAllArtists()
		err := parsedTemplate.Execute(w, Artists)
		if err != nil {
			log.Fatal(err)
		}
	})

	fmt.Println("Server is running on localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
