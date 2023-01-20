/*
Interargit avec l'api donné : https://groupietrackers.herokuapp.com/api
*/

package apiFunctions

import (
	"fmt"
	"net/http"
	"io/ioutil"
	// "encoding/json"
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

func GetConcerts(relations string) []string {
	var concertsData [][]string
	var start, inside int
	for i1, i2 := range relations {
		if i2 == '{' {
			start = i1+1
			inside = 1
		}
		if (i2 == '}') && (inside == 1) {
			concertsData = append(concertsData, FormatConcertString(relations[start: i1+1]))	
			inside = 0
		}
	}

	var data []string
	index := 0
	for _, i := range concertsData {
		for _, k := range i {
			data = append(data, "")
			data[index] += k+"\n"
		}
		index++
	}

	return data
}

func FormatConcertString(s string) []string {
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

func Search(input string, artistsData Artists) int {
	input = Minimalize(input)
	var similarities []int
	id := -1
	for _, i := range artistsData {
		name := Minimalize(i.Name)
		if input == Minimalize(name) {
			id = i.ID
		}
		similarities = append(similarities, CheckSimilarities(input, name))
	}

	max := 0
	maxId := -1
	var maxSimilarities []int
	for len(maxSimilarities) < 5 {
		for i1, i2 := range similarities {
			if i2 > max {
				for _, k := range maxSimilarities {
					if i1 == k {
						break
					} else if i1 == len(maxSimilarities) {
						max = i2
						maxId = i1
					}
				}
			}
		}
		maxSimilarities = append(maxSimilarities, maxId)
	}

	fmt.Println(maxSimilarities)

	return id
}

func Minimalize(s string) string {
	var solu string
	for i, _ := range s {
		if (s[i] >= 'A') && (s[i] <= 'Z') {
			solu += string(s[i]+('a'-'A'))
		} else {
			solu += string(s[i])
		}
	}

	return solu
}

func CheckSimilarities(s1 string, s2 string) int {
	similarities := 0
	for _, k := range s1 {
		for _, j := range s2 {
			if k == j {
				similarities ++
				break
			}
		}
	}

	return similarities
}

func FilterCreationDate(artistsData Artists, mode int, date int) []int {
	var idOverDate []int
	for _, i := range artistsData {
		switch mode {
		case 0:
			if i.CreationDate == date {
				idOverDate = append(idOverDate, i.ID)
			}
		case 1:
			if i.CreationDate < date {
				idOverDate = append(idOverDate, i.ID)
			}
		case 2:
			if i.CreationDate > date {
				idOverDate = append(idOverDate, i.ID)
			}
		}
	}

	return idOverDate
}