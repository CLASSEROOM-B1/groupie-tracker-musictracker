package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type APIrep []struct {
	Id   int    `json:"Id"`
	Name string `json:"Name"`

	//`json:"data"`
}

func GetApi(url string) APIrep {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var response APIrep
	json.Unmarshal(body, &response)
	return response

}

func main() {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	response := GetApi(url)

	for _, data := range response {
		fmt.Println("ID:", data.Id)
		fmt.Println("NAME", data.Name)
	}

}
