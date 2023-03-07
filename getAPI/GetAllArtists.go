package musicTracker

import (
	"encoding/json"
	"log"
	"strconv"
)

func GetAllArtists() []ArtistLight {
	var artistsList []ArtistLight
	err := json.Unmarshal(GetData("groupietrackers.herokuapp.com/api/artists"), &artistsList)
	if err != nil {
		log.Fatal(err)
	}
	for i := range artistsList {
		err = json.Unmarshal(GetData("groupietrackers.herokuapp.com/api/locations/"+strconv.Itoa(i+1)), &artistsList[i].LocationsStruct)
		if err != nil {
			log.Fatal(err)
		}
	}
	return artistsList
}
