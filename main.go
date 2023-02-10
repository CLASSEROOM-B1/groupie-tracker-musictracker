package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var client *http.Client

type deezer struct {
	music        string `json:"music"`
	artist       string `json:"artist"`
	concert      string `json:"concert"`
	localisation string `json:"localisation"`
}

//	type RandomUser struct {
//		Results []UserResult
//	}
//
//	type UserResult struct {
//		Name    UserResult
//		Email   string
//		Picture UserPicture
//	}
type UserName struct {
	Title string
	First string
	Last  string
}
type UserPicture struct {
	Large     string
	Medium    string
	Thumbnail string
}

func Getdeezer() {
	url := "https://api.deezer.com/artist"
	var deezer deezer
	err := GetJson(url, &deezer)
	if err != nil {
		fmt.Printf("error getting deezer: %\n", err.Error())
		return
	}
}

// func GetRandomUser() {
// 	url := "https://randomuser.me/api/?inc=name,email,picture"
// 	var user RandomUser
// 	err := GetJson(url, &user)
// 	if err != nil {
// 		fmt.Printf("error getting deezer fact : %\n", err.Error())
// 	}else {
// 		fmt.Printf("User: %s %s %s\nEmail: %s\nThumbnail:  %s",
// 		&user.Results[0].Name
// }
// }

func GetJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

func mainss() {
	client = &http.Client{Timeout: 10 * time.Second}
	Getdeezer()
	deezer := deezer{
		music:  "bass",
		artist: "NM",
	}
	jsonStr, err := json.Marshal(deezer)
	if err != nil {
		fmt.Printf("error marshaling: %\n", err.Error())
	} else {
		fmt.Printf("Test JSON: %\n", string(jsonStr))
	}

}
