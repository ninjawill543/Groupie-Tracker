/*
Lance le serveur
*/

package main

import (
	"net/http"
	"html/template"
	"fmt"
	api "apiFunctions/api"
	"encoding/json"
	"strconv"
)
	
// Structure pour contenir les données de l'api
type datasJson struct {
	Artists api.Artists // Informations générales sur le groupe
	Locations api.Locations // Lieux des concerts
	Dates api.Dates // Dates des concerts
	Relations []string
}

var groupData datasJson

func main() {
	groupData = InitializeData()
	
	fs := http.FileServer(http.Dir("./static/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("./static/index.html")) // Affiche la page

	// Affiche dans le terminal l'activité sur le site
	switch r.Method {
	case "GET":
		fmt.Println("GET")
	case "POST": // Gestion d'erreur
		if err := r.ParseForm(); err != nil {
			return
		}
	}

	input := r.Form.Get("search") // Récupère la recherche

	// Récupère la date de creation dans les filtres
	creationDateMin := r.Form.Get("creation-date-min")
	creationDateMax := r.Form.Get("creation-date-max")
	// Récupère la date du première album dans les filtres
	firstAlbumDateMin := r.Form.Get("first-album-date-min")
	firstAlbumDateMax := r.Form.Get("first-album-date-max")
	// Récupère une liste avec les nombres de membres
	members := r.Form.Get("members")
	// Récupère la location dans les filtres
	location := r.Form.Get("location")

	fmt.Println(creationDateMin, creationDateMax, firstAlbumDateMin, firstAlbumDateMax, members, location)

	if len(creationDateMin) > 0 {
		mockData := InitializeData()
		var id []int
		for i := 1; i < 52; i++ {
			id = append(id, i)
		}

		min, err1 := strconv.Atoi(creationDateMin)
		max, err2 := strconv.Atoi(creationDateMax)
		if (err1 == nil) && (err2 == nil) {
			id = api.FilterCreationDate(groupData.Artists, min-1, max+1)
			groupData = IdToJson(mockData, id)
		} else {
			fmt.Println("error main creationDate")
			fmt.Println(err1, err2)
		}

		min, err1 = strconv.Atoi(firstAlbumDateMin)
		max, err2 = strconv.Atoi(firstAlbumDateMax)
		if (err1 == nil) && (err2 == nil) {
			id = api.FilterFirstAlbum(groupData.Artists, min-1, max+1)
			groupData = IdToJson(mockData, id)
		} else {
			fmt.Println("error main firstAlbum")
			fmt.Println(err1, err2)
		}

		if len(members) > 0 {
			number, err := strconv.Atoi(members)
			if err == nil {
				id = api.FilterMembers(groupData.Artists, number)
				groupData = IdToJson(mockData, id)
			} else {
				fmt.Println("error main members")
				fmt.Println(err)
			}
		}

		if len(location) > 0 {
			idLocation := api.FilterLocations(groupData.Locations, location)
			var newId []int
			for _, i := range id {
				for _, k := range idLocation {
					if i == k {
						newId = append(newId, i)
					}
				}
			}
			id = newId
			groupData = IdToJson(mockData, newId)
		}

		fmt.Println(id)
	}

	if len(input) > 0 {
		result := api.Search(input, groupData.Artists)
		groupData = IdToJson(groupData, result)
	}

	tmpl.Execute(w, groupData) // Execute le code html en fonction des changements de variables
}

// Récupère les données de l'api et les stockent dans une structure datasJson sous forme de json
func InitializeData() datasJson {
	var apiData datasJson
	// Récupère les données en tableau de bytes et les formata en Json
	json.Unmarshal(api.ExtractRawData(0), &apiData.Artists)
	json.Unmarshal(api.ExtractRawData(1), &apiData.Locations)
	json.Unmarshal(api.ExtractRawData(2), &apiData.Dates)
	apiData.Relations = api.GetConcerts(string(api.ExtractRawData(3)))

	return apiData
}

func IdToJson(groupData datasJson, id []int) datasJson {
	var data datasJson
	for _, i := range id {
		data.Artists = append(data.Artists, groupData.Artists[i])
		data.Locations.Index = append(data.Locations.Index, groupData.Locations.Index[i])
		data.Dates.Index = append(data.Dates.Index, groupData.Dates.Index[i])
		data.Relations = append(data.Relations, groupData.Relations[i])
	}

	return data
}