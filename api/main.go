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
	apiData := initializeData()
	fmt.Println(getDataById(apiData.artists, 1))
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
	json.Unmarshal(extractRawData(apiUrls[0]), &apiData.artists)
	json.Unmarshal(extractRawData(apiUrls[1]), &apiData.locations)
	json.Unmarshal(extractRawData(apiUrls[2]), &apiData.dates)
	json.Unmarshal(extractRawData(apiUrls[3]), &apiData.relations)

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
			groupData = append(groupData, i.Locations)
			groupData = append(groupData, "§")
			groupData = append(groupData, i.ConcertDates)
			groupData = append(groupData, "§")
			groupData = append(groupData, i.Relations)
			break
		}
	}

	return groupData
}

func Concerts