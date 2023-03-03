package musicTracker

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func GetData(lien string) []byte {
	data, err := http.Get("https://" + lien)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(data.Body)
	body, err := ioutil.ReadAll(data.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}
