package main

import (
	"fmt"
	"time"
)

type Band struct {
	Name               string
	Formed             time.Time
	FirstAlbumReleased time.Time
	Members            int
	ConcertLocations   []string
}

func Filters() {
	bands := []Band{
		{
			Name:               "The Beatles",
			Formed:             time.Date(1960, time.February, 1, 0, 0, 0, 0, time.UTC),
			FirstAlbumReleased: time.Date(1963, time.March, 22, 0, 0, 0, 0, time.UTC),
			Members:            4,
			ConcertLocations:   []string{"Liverpool", "London", "New York", "Los Angeles"},
		},
		{
			Name:               "Led Zeppelin",
			Formed:             time.Date(1968, time.October, 12, 0, 0, 0, 0, time.UTC),
			FirstAlbumReleased: time.Date(1969, time.January, 12, 0, 0, 0, 0, time.UTC),
			Members:            4,
			ConcertLocations:   []string{"London", "New York", "Los Angeles"},
		},
		{
			Name:               "Queen",
			Formed:             time.Date(1970, time.July, 13, 0, 0, 0, 0, time.UTC),
			FirstAlbumReleased: time.Date(1973, time.July, 13, 0, 0, 0, 0, time.UTC),
			Members:            4,
			ConcertLocations:   []string{"London", "Berlin", "Paris", "New York"},
		},
	}

	// Filter by formation date
	minFormedDate := time.Date(1965, time.January, 1, 0, 0, 0, 0, time.UTC)
	maxFormedDate := time.Date(1975, time.December, 31, 0, 0, 0, 0, time.UTC)
	formedBands := make([]Band, 0)
	for _, band := range bands {
		if band.Formed.After(minFormedDate) && band.Formed.Before(maxFormedDate) {
			formedBands = append(formedBands, band)
		}
	}
	fmt.Println("Bands formed between", minFormedDate, "and", maxFormedDate)
	for _, band := range formedBands {
		fmt.Println("\t", band.Name)
	}

	// Filter by first album release date
	minFirstAlbumDate := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
	maxFirstAlbumDate := time.Date(1980, time.December, 31, 0, 0, 0, 0, time.UTC)
	firstAlbumBands := make([]Band, 0)
	for _, band := range bands {
		if band.FirstAlbumReleased.After(minFirstAlbumDate) && band.FirstAlbumReleased.Before(maxFirstAlbumDate) {
			firstAlbumBands = append(firstAlbumBands, band)
		}
	}
	fmt.Println("Bands with first album released between", minFirstAlbumDate, "and", maxFirstAlbumDate)
	for _, band := range firstAlbumBands {
		fmt.Println("\t", band.Name)
	}

	// Filter by number of members
	minMembers := 3
	maxMembers := 5
	memberBands := make([]Band, 0)
	for _, band := range bands {
		if band.Members >= minMembers && band.Members <= maxMembers {
			memberBands = append(memberBands, band)
		}
	}
	fmt.Println("Bands with", minMembers, "to", maxMembers, "members")
	for _, band := range memberBands {
		fmt.Println("\t", band.Name)
	}

	// Filter by concert location
	concertLocations := []string{"London", "New York"}
	locationBands := make([]Band, 0)
	for _, band := range bands {
		for _, location := range concertLocations {
			for _, bandLocation := range band.ConcertLocations {
				if bandLocation == location {
					locationBands = append(locationBands, band)
					break
				}
			}
		}
	}
	fmt.Println("Bands with concerts in", concertLocations)
	for _, band := range locationBands {
		fmt.Println("\t", band.Name)
	}
}
