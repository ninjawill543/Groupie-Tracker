/*
Interargit avec l'api donné : https://groupietrackers.herokuapp.com/api/
*/

package apiFunctions

import (
	"fmt"
	"net/http"
	"io/ioutil"
	// "encoding/json"
	"strconv"
)

// Contient les urls de l'api
var apiUrls = [4]string{"https://groupietrackers.herokuapp.com/api/artists", "https://groupietrackers.herokuapp.com/api/locations", "https://groupietrackers.herokuapp.com/api/dates", "https://groupietrackers.herokuapp.com/api/relation"}

// Extrait et retourne les données en bytes de l'url donné
func ExtractRawData(urlIndex int) []byte {
	resp, err := http.Get(apiUrls[urlIndex]) // Récupère le contenu de l'url
	if err != nil { // En cas d'erreur
		fmt.Println("Error loading the API : ", apiUrls[urlIndex])
		panic(err)
	}

	rawData, err := ioutil.ReadAll(resp.Body) // Récupère les données dans l'html
	if err != nil { // En cas d'erreur
		fmt.Println("Error reading the response body : ", apiUrls[urlIndex], "\n", resp)
		panic(err)
	}

	return rawData // Retourne les données en bytes
}

// Récupère l'id d'un nom de groupe donné
func GetId(artistsData Artists, name string) int {
	for _, i := range artistsData {
		if i.Name == name {
			return i.ID
		}
	}
	return -1 // Si aucun groupe avec ce nom n'a été trouvé
}

/* Inutile ?
func GetDataById(artistsData Artists, id int) []string {
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

func GetConcertsById(locationData Locations, dateData Dates, id int) []string {
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
*/