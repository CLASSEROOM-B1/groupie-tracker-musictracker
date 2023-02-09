package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type APIrep struct {
	Data []struct {
		Id   int    `json:"Id"`
		Name string `json:"Name"`
	} //`json:"data"`
}

func GetApi(url string) (APIrep, error) {
	resp, err := http.Get(url)
	if err != nil {
		return APIrep{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return APIrep{}, err
	}

	var response APIrep
	err = json.Unmarshal(body, &response)
	if err != nil {
		return APIrep{}, err
	}

	return response, nil
	// resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(body))
}

func main() {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	response, err := GetApi(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, data := range response.Data {
		fmt.Println("ID:", data.Id)
		fmt.Println("NAME", data.Name)
	}

}
