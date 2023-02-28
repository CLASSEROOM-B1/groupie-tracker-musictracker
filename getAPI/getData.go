package getAPI

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetData(lien string) []byte {
	data, err := http.Get("http://" + lien)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Body.Close()
	body, err := ioutil.ReadAll(data.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}
