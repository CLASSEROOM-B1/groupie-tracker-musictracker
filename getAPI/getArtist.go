package musicTracker

import (
	"encoding/json"
	"log"
	"strconv"
)

func GetArtist(name string) Artist {
	var artistsList []Artist
	err := json.Unmarshal(GetData("groupietrackers.herokuapp.com/api/artists"), &artistsList)
	if err != nil {
		log.Fatal(err)
	}
	var artist Artist
	for i := 0; i < len(artistsList); i++ {
		if artistsList[i].Name == name {
			artist = artistsList[i]
		}
	}
	err = json.Unmarshal(GetData("groupietrackers.herokuapp.com/api/locations/"+strconv.Itoa(artist.Id)), &artist.LocationsStruct)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(GetData("groupietrackers.herokuapp.com/api/dates/"+strconv.Itoa(artist.Id)), &artist.ConcertDatesStruct)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(artist.ConcertDatesStruct.Dates); i++ {
		artist.ConcertDatesStruct.Dates[i] = artist.ConcertDatesStruct.Dates[i][1:]
	}
	err = json.Unmarshal(GetData("groupietrackers.herokuapp.com/api/relation/"+strconv.Itoa(artist.Id)), &artist.RelationsStruct)
	if err != nil {
		log.Fatal(err)
	}
	return artist
}
