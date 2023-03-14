package musicTracker

import (
	"fmt"
	data "musicTracker/src/data"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"text/template"
)

// Création des structures utiles pour l'affichage de la page
type NumberMembers struct {
	Checked string
	Number  int
}

type RangeDate struct {
	Start string
	End   string
}

type PageDataGroups struct {
	Groups            []data.Group
	CheckboxMember    []NumberMembers
	RangeCreationDate RangeDate
	RangeFirstAlbum   RangeDate
	NumberPages       []int
}

// Création de variables utiles au lancement de la page
var nbMembersMax = 8
var creationDate = RangeDate{"1958", "2015"}
var firstAlbum = RangeDate{"1963", "2018"}

// Fonction lancée à /artistes
func Artistes(w http.ResponseWriter, req *http.Request) {

	// importation des templates utiles pour l'affichage de la page
	const path = "./template/artistes.html"
	t, e := template.ParseFiles(path, "./template/layout/header.html")

	// Boucle sur le nombre de checkbox à afficher
	var tabMembers []NumberMembers
	for i := 1; i <= nbMembersMax; i++ {
		var member = NumberMembers{"unchecked", i}
		tabMembers = append(tabMembers, member)
	}

	// Importation de tous les groupes + tri des groupes par nom
	groups := data.GetGroups()
	sort.SliceStable(groups, func(i, j int) bool {
		return groups[i].Name < groups[j].Name
	})

	// Création d'une variable qui stockera les groupes correspondants aux filtres
	var tabGroups []data.Group

	if req.Method == "POST" {
		creationDate = RangeDate{req.FormValue("startCD"), req.FormValue("endCD")}
		firstAlbum = RangeDate{req.FormValue("startFA"), req.FormValue("endFA")}

		fmt.Printf("Années : %s\n", creationDate)
		fmt.Printf("Album : %s\n", firstAlbum)

		// Boucle sur les membres, on voit lesquels sont checked ou pas
		var counterUnchecked = 0
		for j := 1; j <= nbMembersMax; j++ {
			n := strconv.Itoa(j)
			if req.FormValue("checkbox-"+n) == n {
				// On change la checkbox en checked si elle a été cochée
				tabMembers[j-1] = NumberMembers{"checked", j}
			} else {
				counterUnchecked += 1
			}
		}
		fmt.Printf("Membres : %v\n", tabMembers)

		// TRI PAR MEMBRES
		if counterUnchecked != 8 {
			// On va chercher le nombre de membres par groupe
			for k := range groups {
				lenGroupMembers := len(groups[k].Members)
				for l := range tabMembers {
					// On ajoute le groupe si le nombre de membres correspond
					if lenGroupMembers == tabMembers[l].Number && tabMembers[l].Checked == "checked" {
						tabGroups = append(tabGroups, groups[k])
					}
				}
			}
		} else {
			// Si aucune checkbox n'a été cochée, le tableau des groupes reste inchangé
			tabGroups = groups
		}
		// TRI PAR MEMBRES FINI

		// TRI PAR CREATION DATE
		yearStartCD, _ := strconv.Atoi(creationDate.Start)
		yearEndCD, _ := strconv.Atoi(creationDate.End)
		// Création d'un nouveau tableau de groupes par rapport à l'année de création
		var tabGroups2 []data.Group
		// On boucle sur le tableau (déjà trié par membres)
		for _, v := range tabGroups {
			if v.CreationDate >= yearStartCD && v.CreationDate <= yearEndCD {
				// On ajoute les groupes dont la Creation Date correspond à l'intervalle
				tabGroups2 = append(tabGroups2, v)
			}
		}
		tabGroups = tabGroups2
		//CREATION DATE FINI

		// TRI PAR FIRST ALBUM
		yearStartFA, _ := strconv.Atoi(firstAlbum.Start)
		yearEndFA, _ := strconv.Atoi(firstAlbum.End)
		// Création d'un nouveau tableau de groupes par rapport à l'année du premier album
		var tabGroups3 []data.Group

		// On boucle sur le tableau (déjà trié par membres et par Creation Date)
		for _, v := range tabGroups {
			// Création d'une Regexp pour récupérer l'année du First Album
			var re, _ = regexp.Compile(`....$`)
			FirstAlbumYear := re.FindStringSubmatch(v.FirstAlbum)[0]
			FirstAlbumYearInt, _ := strconv.Atoi(FirstAlbumYear)

			if FirstAlbumYearInt >= yearStartFA && FirstAlbumYearInt <= yearEndFA {
				// On ajoute les groupes dont l'année de First Album () correspond à l'intervalle
				tabGroups3 = append(tabGroups3, v)
			}
		}
		tabGroups = tabGroups3
		// FIRST ALBUM FINI

	} else {
		// Si aucun critères de tri n'a été sélectionné, on récupère tous les groupes
		tabGroups = groups
	}

	// On récupère un tableau contenant les numéros des pages en fonction du nombre d'éléments
	tabNumberPages := Paging(tabGroups)

	// On calcule l'index du premier et du dernier groupe à afficher
	// En fonction de la page cliquée
	var startElem int
	var endElem int

	page, _ := strconv.Atoi(req.FormValue("page"))
	startElem = (page - 1) * 9
	endElem = startElem + 9

	if endElem > len(tabGroups) {
		endElem = startElem + (len(tabGroups) - startElem)
	}
	if page == 0 {
		startElem = 0
		endElem = 9
	}
	// Fin du calcul des index

	// Regroupement de toutes les données nécessaires pour l'affichage de la page
	tabGroups = tabGroups[startElem:endElem]

	pageGroups := PageDataGroups{
		Groups:            tabGroups,
		CheckboxMember:    tabMembers,
		RangeCreationDate: creationDate,
		RangeFirstAlbum:   firstAlbum,
		NumberPages:       tabNumberPages,
	}

	//gestion de l'erreur 500
	if e != nil {
		http.Error(w, "Internal Serveur Error 500", http.StatusInternalServerError)
		return
	} else {
		//Exécution de la page avec les données de pageGroups
		t.Execute(w, pageGroups)
		fmt.Print("Artistes")
	}

}

func Paging(tabGroups []data.Group) []int {
	numberOfPages := len(tabGroups) / 9
	if len(tabGroups)%9 != 0 {
		numberOfPages++
	}
	var tabNumbers []int
	for i := 1; i <= numberOfPages; i++ {
		tabNumbers = append(tabNumbers, i)
	}
	return tabNumbers
}
