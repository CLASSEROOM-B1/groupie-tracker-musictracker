package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Personnage struct {
	character string
	episode   string
}

func mains() {

	resp, err := http.Get("https://rickandmortyapi.com/api/character/17")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
