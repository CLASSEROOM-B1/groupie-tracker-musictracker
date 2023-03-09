package musicTracker

import ( //Importation des librairies nécessaire
	"fmt"
	"net/http"
	"text/template"
)

func Accueil(w http.ResponseWriter, req *http.Request) { //fonction de qui affiche la page Accueil

	const path = "./template/home.html"

	//importation des page HTML
	t, e := template.ParseFiles(path, "./template/layout/header.html")

	//gestion de l'erreur 500
	if e != nil {
		http.Error(w, "Internal Serveur Error 500", http.StatusInternalServerError)
		return
	}

	//gestion de l'erreur 404

	if req.URL.Path == "/" { //verification de l'URL
	} else if req.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	t.Execute(w, nil)
	fmt.Print("Home")
}
