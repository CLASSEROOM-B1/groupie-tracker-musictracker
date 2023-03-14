package main

//Librairies utilisées
import (
	"fmt"
	handlers "musicTracker/src"
	"net/http"
)

func main() {
	fmt.Print("Server started... 🎸\n")

	// Récupération des fichiers static pour l'affichage des pages
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	//référencement des pages html
	http.HandleFunc("/", handlers.Accueil)
	http.HandleFunc("/artistes", handlers.Artistes)
	http.HandleFunc("/locations", handlers.Locations)
	http.HandleFunc("/events", handlers.Events)
	http.HandleFunc("/cities", handlers.ConcertsLocations)

	//Port du serveur
	http.ListenAndServe(":8080", nil)

}
