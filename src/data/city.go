package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type CityStruct struct {
	City []string `json:"locations"` // "isoler" les locations dans notre structure
}

func GetCity(thisID string) []string {

	url := "https://groupietrackers.herokuapp.com/api/locations/" + thisID // récupération de l'url

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}
	// Gestion d'erreur
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Page-Locations", "All-City")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	// Unmarshal les données sous forme de structure
	thisCity := CityStruct{}
	jsonErr := json.Unmarshal(body, &thisCity)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return thisCity.City // return le tableau de string
}
