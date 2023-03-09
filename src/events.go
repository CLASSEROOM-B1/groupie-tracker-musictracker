package musicTracker

import (
	"fmt"
	"net/http"
	"regexp"
	"sort"
	"text/template"

	data "musicTracker/src/data"
)

// Création de 2 structures contenant les données à envoyer sur la page
type PageDataEvents struct {
	Groups    []data.Group
	Concert   []data.Concert
	ThisGroup data.OneGroup
}
type PageDataEvents2 struct {
	Groups      []data.Group
	GroupsFound []data.Group
	ValueSearch string
	ThisGroup   data.OneGroup
	Concert     []data.Concert
}

// Fonction lancée à /events
func Events(w http.ResponseWriter, req *http.Request) {

	//Importaiton des templates utiles pour l'affichage de la page
	const path = "./template/events.html"
	t, e := template.ParseFiles(path, "./template/layout/header.html")

	// Importation de tous les groupes
	groups := data.GetGroups()
	sort.SliceStable(groups, func(i, j int) bool {
		return groups[i].Name < groups[j].Name
	})

	// Création de 2 variables
	var tabGroupsChecked []data.Group // contiendra les groupes correspondants à la recherche
	var valueSearch string            // contiendra la recherche

	if req.Method == "POST" {
		// Récupération de la recherche
		search := req.FormValue("search")

		// Boucle sur tous les groupes
		for _, v := range groups {

			// Création d'une Regexp pour vérifier si le gorupe correspond
			checkName, _ := regexp.MatchString("(?i)"+search, v.Name)
			// En fonction du nom du groupe
			if checkName {
				// Ajout du groupe si correspond
				tabGroupsChecked = append(tabGroupsChecked, v)
				continue
			} else {
				// En fonction des noms des membres
				for _, member := range v.Members {
					checkMember, _ := regexp.MatchString("(?i)"+search, member)
					if checkMember {
						// Ajout du groupe si correspond
						tabGroupsChecked = append(tabGroupsChecked, v)
						break
					}
				}
			}
		}
		valueSearch = search
	} else {
		tabGroupsChecked = nil
		valueSearch = ""
	}

	// Création de variables pour récupérer le nom et les concerts d'un groupe
	var thisGroup data.OneGroup
	var groupConcert []data.Concert
	if req.FormValue("oneGroup") == "" {
		groupConcert = nil

	} else {
		// Récupération des données du groupe en question
		thisGroup = data.GetOneGroup(req.FormValue("oneGroup"))
		groupConcert = data.GetEvents(req.FormValue("oneGroup"))
	}

	// Regroupement de toutes les données nécessaires pour l'affichage de la page
	pageEvents := PageDataEvents2{
		Groups:      groups,
		GroupsFound: tabGroupsChecked,
		ValueSearch: valueSearch,
		ThisGroup:   thisGroup,
		Concert:     groupConcert,
	}

	//gestion de l'erreur 500
	if e != nil {
		http.Error(w, "Internal Serveur Error 500", http.StatusInternalServerError)
		return
	} else {
		//Exécution de la page avec les données de pageGroups
		t.Execute(w, pageEvents)
		fmt.Print("Events ")
	}

}
