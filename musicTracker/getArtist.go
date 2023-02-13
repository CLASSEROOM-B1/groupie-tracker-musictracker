package musicTracker

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Artist struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    Locations
	ConcertDates ConcertDates
	Relations    Relations
}

type Locations struct {
	Locations []string
}

type ConcertDates struct {
	Dates []string
}

type Relations struct {
	DatesLocations map[string][]string
}

func GetArtist(name string) Artist {
	var artistsList []Artist
	json.Unmarshal(GetData("groupietrackers.herokuapp.com/api/artists"), &artistsList)
	var artist Artist
	for i := 0; i < len(artistsList); i++ {
		if artistsList[i].Name == name {
			artist = artistsList[i]
		}
	}
	json.Unmarshal(GetData("groupietrackers.herokuapp.com/api/locations/"+strconv.Itoa(artist.Id)), &artist.Locations)
	json.Unmarshal(GetData("groupietrackers.herokuapp.com/api/dates/"+strconv.Itoa(artist.Id)), &artist.ConcertDates)
	for i := 0; i < len(artist.ConcertDates.Dates); i++ {
		artist.ConcertDates.Dates[i] = artist.ConcertDates.Dates[i][1:]
	}
	json.Unmarshal(GetData("groupietrackers.herokuapp.com/api/relation/"+strconv.Itoa(artist.Id)), &artist.Relations)
	fmt.Println(artist)
	return artist
}
