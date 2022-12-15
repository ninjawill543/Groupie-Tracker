package main

import (
	apiStructures "structures/structures"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

var artistsUrl = "https://groupietrackers.herokuapp.com/api/artists"
var locationsUrl = "https://groupietrackers.herokuapp.com/api/locations"
var datesUrl = "https://groupietrackers.herokuapp.com/api/dates"
var relationsUrl = "https://groupietrackers.herokuapp.com/api/relation"

func main() {
	var rawData []byte
	
	rawData = extractData(artistsUrl)
	artistsData := apiStructures.Artists{}
	json.Unmarshal(rawData, &artistsData)

	rawData = extractData(locationsUrl)
	locationData := apiStructures.Locations{}
	json.Unmarshal(rawData, &locationData)

	rawData = extractData(datesUrl)
	datesData := apiStructures.Dates{}
	json.Unmarshal(rawData, &datesData)

	rawData = extractData(relationsUrl)
	relationsData := apiStructures.Relations{}
	json.Unmarshal(rawData, &relationsData)
}

func extractData (url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error loading the API.")
		panic(err)
	}

	rawData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response body.")
		panic(err)
	}

	return rawData
}
