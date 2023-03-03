package musicTracker

import (
	"encoding/json"
	"fmt"
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
	fmt.Println(artistsList)
	return artistsList
}
