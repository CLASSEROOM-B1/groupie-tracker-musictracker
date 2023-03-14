package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// Création d'une structure pour stocker les relations d'un groupe
type Infos struct {
	ID        int         `json:"id"`
	TabEvents interface{} `json:"datesLocations"`
}

// Création d'une structure pour stocker les données du tableau d'évènements du groupe
type Concert struct {
	City    string
	Country string
	Dates   []string
}

func GetEvents(thisID string) []Concert {
	url := "https://groupietrackers.herokuapp.com/api/relation/" + thisID
	fmt.Print("Data Concert OK")
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Page-Events", "All-Dates")

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

	infosTab := Infos{}
	// On Unmarshal les données trouvées sous forme de structure
	jsonErr := json.Unmarshal(body, &infosTab)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	// Une des valeurs de la structure est une interface. On travaille dessus pour la structurer
	// Création d'une variable qui parcourt le tableau d'interfaces et qui créé une map.
	mapInfos := infosTab.TabEvents.(map[string]interface{})

	// Création d'un tableau de structures avec les villes et les dates.
	var tabConcert []Concert

	// On y ajoute la clé et la valeur de chaque évènement.
	for k, v := range mapInfos {
		// On sépare la ville et le pays
		tabLocation := strings.Split(k, "-")
		city := strings.Title(strings.Replace(tabLocation[0], "_", " ", -1))
		country := strings.Title(strings.Replace(tabLocation[1], "_", " ", -1))

		// On ajoute chaque date
		var tabValues []string
		for _, v2 := range v.([]interface{}) {
			tabValues = append(tabValues, v2.(string))
		}

		// tableau de Concert dans lequel on ajoute tout
		var oneConcert = Concert{city, country, tabValues}
		tabConcert = append(tabConcert, oneConcert)
	}

	// On retourne le tableau de Concert qui contient les objets correspondants aux concerts
	return tabConcert
}
