package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Création d'une structure pour stocker les données d'un groupe
// On utilise les clés de l'API pour permettre le Marshalling
type OneGroup struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

// création d'une fonction qui va chercher les infos du groupe concerné
func GetOneGroup(thisID string) OneGroup {
	url := "https://groupietrackers.herokuapp.com/api/artists/" + thisID
	fmt.Print("Lancement Data ONE GROUP OK ")

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	// On va chercher les données à l'url demandée
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Page-Events", "One-Group")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	// On lit les données
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	// On récupère les données puis on les Unmarshal grâce à la structure créée.
	thisGroup := OneGroup{}
	jsonErr := json.Unmarshal(body, &thisGroup)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	// On retourne un object de structure OneGroup{}

	return thisGroup
}
