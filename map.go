package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Location is a struct to store the latitude and longitude of a concert location
type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func localisation() {
	// Your API key for the Google Maps Geocoding API
	apiKey := "AIzaSyD4OnjwP-3hpqykIROvV96HobEWtHJPxVA"

	// The list of concert locations for a specific artist or band
	concertLocations := []string{"Mayence, Allemagne", "Paris, France", "Londres, Royaume-Uni"}

	// Make a map to store the converted location coordinates
	locations := make(map[string]Location)

	// Loop through each concert location
	for _, location := range concertLocations {
		// Replace spaces with plus signs for the API request
		location = strings.Replace(location, " ", "+", -1)

		// Make the API request
		response, err := http.Get("https://maps.googleapis.com/maps/api/geocode/json?address=" + location + "&key=" + apiKey)
		if err != nil {
			fmt.Println("Error making API request:", err)
			continue
		}
		defer response.Body.Close()

		// Read the API response
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading API response:", err)
			continue
		}

		// Unmarshal the JSON response into a map
		var data map[string]interface{}
		err = json.Unmarshal(body, &data)
		if err != nil {
			fmt.Println("Error unmarshaling JSON:", err)
			continue
		}

		// Extract the latitude and longitude from the API response
		lat := data["results"].([]interface{})[0].(map[string]interface{})["geometry"].(map[string]interface{})["location"].(map[string]interface{})["lat"].(float64)
		lng := data["results"].([]interface{})[0].(map[string]interface{})["geometry"].(map[string]interface{})["location"].(map[string]interface{})["lng"].(float64)

		// Store the location coordinates in the map
		locations[location] = Location{
			Lat: lat,
			Lng: lng,
		}
	}

	// Print the map of location coordinates
	fmt.Println(locations)
}
