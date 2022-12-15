package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type artists []struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type locations struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

type dates struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

var artistsUrl = "https://groupietrackers.herokuapp.com/api/artists"
var locationsUrl = "https://groupietrackers.herokuapp.com/api/locations"
var datesUrl = "https://groupietrackers.herokuapp.com/api/dates"
var relationUrl = "https://groupietrackers.herokuapp.com/api/relation"

func main() {
	url := locationsUrl

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error loading the API.")
		panic(err)
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response body.")
		panic(err)
	}

	
	// fmt.Println(string(respData))
 	/*
	var m map[string]interface{}
	json.Unmarshal(respData, &m)

	fmt.Println(m)
	*/

	
	data := locations{}
	json.Unmarshal(respData, &data)

	fmt.Println(data)

	
	for _, i := range data.Index {
		fmt.Println(i.Locations)
	}
	
}


