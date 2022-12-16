package main

import (
	"fmt"
	"encoding/json"
	api "apiFunctions/api"
	//"strconv"
)

// Structure pour contenir les données de l'api
type datasJson struct {
	artists api.Artists // Informations générales sur le groupe
	locations api.Locations // Lieux des concerts
	dates api.Dates // Dates des concerts
	relations string
}

func main() {
	data := InitializeData()
	concertsData := GetConcerts(data.relations)

	for _, i := range concertsData {
		for _, k := range i {
			fmt.Println(k)
		}
	}
}

// Récupère les données de l'api et les stockent dans une structure datasJson sous forme de json
func InitializeData() datasJson {
	var apiData datasJson
	// Récupère les données en tableau de bytes et les formata en Json
	json.Unmarshal(api.ExtractRawData(0), &apiData.artists)
	json.Unmarshal(api.ExtractRawData(1), &apiData.locations)
	json.Unmarshal(api.ExtractRawData(2), &apiData.dates)
	apiData.relations = string(api.ExtractRawData(3))

	return apiData
}

func GetConcerts(relations string) [][]string {
	fmt.Println(relations, "\n")

	var concertsData [][]string
	var start, inside int
	for i1, i2 := range relations {
		if i2 == '{' {
			start = i1+1
			inside = 1
		}
		if (i2 == '}') && (inside == 1) {
			concertsData = append(concertsData, formatConcertString(relations[start: i1+1]))	
			inside = 0
		}
	}

	return concertsData
}

func formatConcertString(s string) []string {
	var formated []string
	var index, start, iteration, length int = 0, -1, 0, 0
	for i1, i2 := range s {
		if (i2 == ']') && (iteration != 0) {
			index++
			iteration = 0
		}
		if string(i2) == "\"" {
			if start == -1 {
				start = i1+1
			} else {
				switch iteration {
				case 0:
					formated = append(formated, s[start: i1])
					length = len(formated[index])+4
				case 1:
					for i := 0; i < 4; i++ {
						formated[index] += " "
					}
					formated[index] += s[start: i1]
				default:
					formated[index] = formated[index]+"\n"
					for i := 0; i < length; i++ {
						formated[index] += " "
					}
					formated[index] += s[start: i1]
				}
				start = -1
				iteration++
			}
		}
	}

	return formated
}