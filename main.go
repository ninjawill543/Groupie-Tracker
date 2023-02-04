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
	creationDate := r.Form.Get("creation-date")
	// Récupère la date du première album dans les filtres
	firstAlbumDate := r.Form.Get("first-album-date")
	// Récupère une liste avec les nombres de membres
	members := r.Form.Get("members")
	// Récupère la location dans les filtres
	location := r.Form.Get("location")

	fmt.Println(creationDate, firstAlbumDate, members, location)

	if len(input) > 0 {
		fmt.Println("Request : ", input)
		result := api.Search(input, groupData.Artists)
		fmt.Println(IdToJson(groupData, result))
		groupData = IdToJson(groupData, result)
	} else {
		groupData = InitializeData()
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
		data.Relations = append(data.Relations, groupData.Relations[i])
	}
	return data
}