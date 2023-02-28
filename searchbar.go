package main

import (
	"fmt"
	"strings"
)

type Artist struct {
	Name         string
	Members      []string
	FirstAlbum   string
	CreationDate string
}

func (a Artist) String() string {
	return a.Name
}

func searchArtists(query string, artists []Artist) []Artist {
	query = strings.ToLower(query)
	results := []Artist{}
	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), query) ||
			contains(artist.Members, query) ||
			strings.Contains(strings.ToLower(artist.FirstAlbum), query) ||
			strings.Contains(strings.ToLower(artist.CreationDate), query) {
			results = append(results, artist)
		}
	}
	return results
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if strings.ToLower(a) == strings.ToLower(e) {
			return true
		}
	}
	return false
}

func artist() {
	artists := []Artist{
		{
			Name:         "Queen",
			Members:      []string{"Freddie Mercury", "John Deacon", "Brian May", "Roger Taylor"},
			FirstAlbum:   "Queen",
			CreationDate: "1970",
		},
		{
			Name:         "The Beatles",
			Members:      []string{"John Lennon", "Paul McCartney", "George Harrison", "Ringo Starr"},
			FirstAlbum:   "Please Please Me",
			CreationDate: "1960",
		},
	}
	results := searchArtists("queen", artists)
	fmt.Println(results)
}
