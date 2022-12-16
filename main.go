/*
Lance le serveur
*/

package main

import (
	"net/http"
	// "html/template"
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

func main() {
	data := InitializeData()
	fmt.Println(data.artists)

	fs := http.FileServer(http.Dir("./static/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {

	// tmpl := template.Must(template.ParseFiles("./static/index.html")) // Affiche la page

	// Affiche dans le terminal l'activité sur le site
	switch r.Method {
	case "GET":
		fmt.Println("GET")
	case "POST": // Gestion d'erreur
		if err := r.ParseForm(); err != nil {
			return
		}
	}

	/*
	// variable := r.Form.Get("input") // Récupère la saisie du joueur

	// Transmet au code html les variables nécessaires
	
	data := Data{
	}

	tmpl.Execute(w, data) // Execute le code html en fonction des changements de variables
	*/
}

// Récupère les données de l'api et les stockent dans une structure datasJson sous forme de json
func InitializeData() datasJson {
	var apiData api.datasJson
	// Récupère les données en tableau de bytes et les formata en Json
	json.Unmarshal(api.ExtractRawData(0), &apiData.artists)
	json.Unmarshal(api.ExtractRawData(1), &apiData.locations)
	json.Unmarshal(api.ExtractRawData(2), &apiData.dates)
	// json.Unmarshal(ExtractRawData(apiUrls[3]), &apiData.relations)

	return apiData
}