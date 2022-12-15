package main

import (
	apiStructures "structures/structures"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

var apiUrls = [4]string{"https://groupietrackers.herokuapp.com/api/artists", "https://groupietrackers.herokuapp.com/api/locations", "https://groupietrackers.herokuapp.com/api/dates", "https://groupietrackers.herokuapp.com/api/relation"}

type datasJson struct {
	artists apiStructures.Artists
	locations apiStructures.Locations
	dates apiStructures.Dates
	relations apiStructures.Relations
}

func main () {
	apiData := initializeData()
	fmt.Println(apiData.artists)
}

func extractRawData (url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error loading the API : ",url)
		panic(err)
	}

	rawData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response body : ",url,"\n",resp)
		panic(err)
	}

	return rawData
}

func initializeData () datasJson {
	var apiData datasJson
	json.Unmarshal(extractRawData(apiUrls[0]), &apiData.artists)
	json.Unmarshal(extractRawData(apiUrls[1]), &apiData.locations)
	json.Unmarshal(extractRawData(apiUrls[2]), &apiData.dates)
	json.Unmarshal(extractRawData(apiUrls[3]), &apiData.relations)

	return apiData
}