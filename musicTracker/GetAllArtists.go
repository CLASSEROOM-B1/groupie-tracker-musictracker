package musicTracker

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func GetAllArtists() []Artist {
	var artistsList []Artist
	json.Unmarshal(GetData("groupietrackers.herokuapp.com/api/artists"), &artistsList)
	for i := 0; i < len(artistsList); i++ {
		json.Unmarshal(GetData("groupietrackers.herokuapp.com/api/locations/"+strconv.Itoa(i+1)), &artistsList[i].Locations)
		json.Unmarshal(GetData("groupietrackers.herokuapp.com/api/dates/"+strconv.Itoa(i+1)), &artistsList[i].ConcertDates)
		for j := 0; j < len(artistsList[i].ConcertDates.Dates); j++ {
			artistsList[i].ConcertDates.Dates[j] = artistsList[i].ConcertDates.Dates[j][1:]
		}
		json.Unmarshal(GetData("groupietrackers.herokuapp.com/api/relation/"+strconv.Itoa(i+1)), &artistsList[i].Relations)
	}
	fmt.Println(artistsList)
	return artistsList
}
