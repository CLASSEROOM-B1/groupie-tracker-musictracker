package musicTracker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	data "musicTracker/src/data"
)

func Locations(w http.ResponseWriter, req *http.Request) {

	groups := data.GetGroups() // importation des groupes

	const path = "./template/locations.html"

	t, e := template.ParseFiles(path, "./template/layout/header.html")

	if req.Method == "POST" {

		fmt.Print("req OK", "\n")

	}

	type DataCity struct {
		Groups []data.Group
	}
	// tableau des location dans lequel on ajoute notre structure DataCity.
	pageLocations := DataCity{Groups: groups}

	//gestion de l'erreur 500
	if e != nil {
		http.Error(w, "Internal Serveur Error 500", http.StatusInternalServerError)
		return
	} else {
		t.Execute(w, pageLocations)
		fmt.Print("Locations")

	}
}

func ConcertsLocations(w http.ResponseWriter, req *http.Request) {

	indexGroup := req.FormValue("groups") // récupèration de l'index de ce que nous retourne la valeur de "groups"
	cities := data.GetCity(indexGroup)    // récupèration grâce à la fonction getCity la valeurs des villes grâce a l'index

	jsonData, _ := json.Marshal(cities) // mettre ses données sous formats JSON

	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
}
