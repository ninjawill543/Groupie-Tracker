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

	data := apiStructures.Locations{}
	json.Unmarshal(respData, &data)

	fmt.Println(data)

	for _, i := range data.Index {
		fmt.Println(i.Locations)
	}
}


