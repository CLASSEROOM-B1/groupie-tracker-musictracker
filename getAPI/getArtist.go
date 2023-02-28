package musicTracker

import (
	"encoding/json"
	"fmt"
	"log"
	"musicTracker"
	"strconv"
)

func GetArtist(name string) musicTracker.Artist {
	var artistsList []musicTracker.Artist
	err := json.Unmarshal(GetData("groupietrackers.herokuapp.com/api/artists"), &artistsList)
	if err != nil {
		log.Fatal(err)
	}
	var artist musicTracker.Artist
	for i := 0; i < len(artistsList); i++ {
		if artistsList[i].Name == name {
			artist = artistsList[i]
		}
	}
	err = json.Unmarshal(GetData("groupietrackers.herokuapp.com/api/locations/"+strconv.Itoa(artist.Id)), &artist.Locations)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(GetData("groupietrackers.herokuapp.com/api/dates/"+strconv.Itoa(artist.Id)), &artist.ConcertDates)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(artist.ConcertDates.Dates); i++ {
		artist.ConcertDates.Dates[i] = artist.ConcertDates.Dates[i][1:]
	}
	err = json.Unmarshal(GetData("groupietrackers.herokuapp.com/api/relation/"+strconv.Itoa(artist.Id)), &artist.Relations)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(artist)
	return artist
}
