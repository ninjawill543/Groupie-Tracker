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
	relations string
}

var t datasJson

func main() {
	data := InitializeData()
	t = data

	concertsData := api.GetConcerts(data.relations)
	for _, i := range concertsData {
		for _, k := range i {
			fmt.Println(k)
		}
	}
	
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

	variable := r.Form.Get("input") // Récupère la saisie du joueur

	if variable == "test" {
		fmt.Println("test")
	}

	data := datasJson{
		Artists: t.Artists,
		Locations: t.Locations,
		Dates: t.Dates,
	}

	// Transmet au code html les variables nécessaires
	fmt.Println(data.Artists)

	tmpl.Execute(w, data) // Execute le code html en fonction des changements de variables
}

// Récupère les données de l'api et les stockent dans une structure datasJson sous forme de json
func InitializeData() datasJson {
	var apiData datasJson
	// Récupère les données en tableau de bytes et les formata en Json
	json.Unmarshal(api.ExtractRawData(0), &apiData.Artists)
	json.Unmarshal(api.ExtractRawData(1), &apiData.Locations)
	json.Unmarshal(api.ExtractRawData(2), &apiData.Dates)
	apiData.relations = string(api.ExtractRawData(3))

	return apiData
}