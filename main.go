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
	artists api.Artists // Informations générales sur le groupe
	locations api.Locations // Lieux des concerts
	dates api.Dates // Dates des concerts
	// relations Relations
}

var t datasJson

func main() {
	data := InitializeData()
	t = data

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
		artists: t.artists,
		locations: t.locations,
		dates: t.dates,
	}

	// Transmet au code html les variables nécessaires
	fmt.Println(data.artists)

	tmpl.Execute(w, data) // Execute le code html en fonction des changements de variables
}

// Récupère les données de l'api et les stockent dans une structure datasJson sous forme de json
func InitializeData() datasJson {
	var apiData datasJson
	// Récupère les données en tableau de bytes et les formata en Json
	json.Unmarshal(api.ExtractRawData(0), &apiData.artists)
	json.Unmarshal(api.ExtractRawData(1), &apiData.locations)
	json.Unmarshal(api.ExtractRawData(2), &apiData.dates)
	// json.Unmarshal(ExtractRawData(apiUrls[3]), &apiData.relations)

	return apiData
}