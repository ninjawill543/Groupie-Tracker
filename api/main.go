package main

import (
	apiStructures "structures/structures"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strconv"
)

var apiUrls = [4]string{"https://groupietrackers.herokuapp.com/api/artists", "https://groupietrackers.herokuapp.com/api/locations", "https://groupietrackers.herokuapp.com/api/dates", "https://groupietrackers.herokuapp.com/api/relation"}

type datasJson struct {
	artists apiStructures.Artists
	locations apiStructures.Locations
	dates apiStructures.Dates
	relations apiStructures.Relations
}

func main() {
	apiData := InitializeData()
	fmt.Println(GetDataById(apiData.artists, 1))
	fmt.Println(GetConcertsById(apiData.locations, apiData.dates, 1))
}

func ExtractRawData(url string) []byte {
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

func InitializeData() datasJson {
	var apiData datasJson
	json.Unmarshal(ExtractRawData(apiUrls[0]), &apiData.artists)
	json.Unmarshal(ExtractRawData(apiUrls[1]), &apiData.locations)
	json.Unmarshal(ExtractRawData(apiUrls[2]), &apiData.dates)
	json.Unmarshal(ExtractRawData(apiUrls[3]), &apiData.relations)

	return apiData
}

func GetId(artistsData apiStructures.Artists, name string) int {
	for _, i := range artistsData {
		if i.Name == name {
			return i.ID
		}
	}
	return -1
}

func GetDataById(artistsData apiStructures.Artists, id int) []string {
	var groupData []string
	for _, i := range artistsData {
		if (id == i.ID) {
			groupData = append(groupData, strconv.Itoa(i.ID))
			groupData = append(groupData, "§")
			groupData = append(groupData, i.Image)
			groupData = append(groupData, "§")
			groupData = append(groupData, i.Name)
			groupData = append(groupData, "§")
			for _, k := range i.Members {
				groupData = append(groupData, k)
			}
			groupData = append(groupData, "§")
			groupData = append(groupData, strconv.Itoa(i.CreationDate))
			groupData = append(groupData, "§")
			groupData = append(groupData, i.FirstAlbum)
			groupData = append(groupData, "§")
			break
		}
	}

	return groupData
}

func GetConcertsById(locationData apiStructures.Locations, dateData apiStructures.Dates, id int) []string {
	var concertsData []string
	for _, i := range locationData.Index {
		if id == i.ID {
			for _, k := range i.Locations {
				concertsData = append(concertsData, k)
				concertsData = append(concertsData, " ")
				concertsData = append(concertsData, "§")
			}
		}
	}
	index := 1;
	for _, i := range dateData.Index {
		if id == i.ID {
			for _, k := range i.Dates {
				concertsData[index] = k
				index += 3
			}
		}
	}

	return concertsData
}