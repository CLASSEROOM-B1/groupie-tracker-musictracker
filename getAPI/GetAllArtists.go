package musicTracker

import (
	"encoding/json"
	"log"
	"musicTracker"
)

func GetAllArtists() []musicTracker.ArtistLight {
	var artistsList []musicTracker.ArtistLight
	err := json.Unmarshal(GetData("groupietrackers.herokuapp.com/api/artists"), &artistsList)
	if err != nil {
		log.Fatal(err)
	}
	return artistsList
}
