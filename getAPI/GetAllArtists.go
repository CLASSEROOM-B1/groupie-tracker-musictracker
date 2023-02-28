package getAPI

import (
	"encoding/json"
	"log"
)

func GetAllArtists() []ArtistLight {
	var artistsList []ArtistLight
	err := json.Unmarshal(GetData("groupietrackers.herokuapp.com/api/artists"), &artistsList)
	if err != nil {
		log.Fatal(err)
	}
	return artistsList
}
