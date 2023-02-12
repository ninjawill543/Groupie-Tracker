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
	Relations []string // Contient les dates et lieux des concerts
	Maps []string // Contient les liens vers des cartes géolocalisants les concerts
}

var groupData datasJson // Contient toutes les données des groupes affichés sur le site
var maps []string // Contient toutes les maps des groupes (chargé une fois au démarrage pour éviter trop d'attentes)

func main() {
	groupData = InitializeData() // Initialise les données
	maps = api.SendMaps(groupData.Locations) // Cherche les cartes pour tous les groupes
	groupData.Maps = maps // Initialise les cartes

	fmt.Println("\nData loaded") // Signale que les données sont chargées et que le site va se lancer

	fs := http.FileServer(http.Dir("./static/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	groupData = InitializeData() // Réinitialise les données

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

	// Récupère la date de creation à filtrer
	creationDateMin := r.Form.Get("creation-date-min")
	creationDateMax := r.Form.Get("creation-date-max")
	// Récupère la date du première album à filtrer
	firstAlbumDateMin := r.Form.Get("first-album-date-min")
	firstAlbumDateMax := r.Form.Get("first-album-date-max")
	// Récupère les nombres de membres à filtrer
	oneMember := r.Form.Get("one-member")
	twoMembers := r.Form.Get("two-members")
	threeMembers := r.Form.Get("three-members")
	fourMembers := r.Form.Get("four-members")
	fiveMembers := r.Form.Get("five-members")
	sixMembers := r.Form.Get("six-members")
	sevenMembers := r.Form.Get("seven-members")
	// Récupère les lieux de concerts à filtrer
	location := r.Form.Get("location")

	// Affiche les filtres utiliser dans le terminal
	fmt.Println(creationDateMin, creationDateMax, firstAlbumDateMin, firstAlbumDateMax, location, oneMember, twoMember, threeMembers, fourMembers, fiveMembers, sixMembers, sevenMembers )

	// Filtre les données
	if len(creationDateMin) > 0 {
		mockData := groupData // Contient les données de tous les groupes
		var id []int // Contient les ids des groupes filtrés
		for i := 1; i < 52; i++ { // Initialise id
			id = append(id, i)
		}

		// Convertie les dates de création de groupe de strings en int
		min, err1 := strconv.Atoi(creationDateMin)
		max, err2 := strconv.Atoi(creationDateMax)
		if (err1 == nil) && (err2 == nil) {
			id = api.FilterCreationDate(groupData.Artists, min-1, max+1) // Filtre les groupes en fonction
			groupData = IdToJson(mockData, id) // Recalcule groupData en fonction
		} else { // En cas d'erreur
			fmt.Println("error main creationDate")
			fmt.Println(err1, err2)
		}

		// Convertie les dates de premier album de strings en int
		min, err1 = strconv.Atoi(firstAlbumDateMin)
		max, err2 = strconv.Atoi(firstAlbumDateMax)
		if (err1 == nil) && (err2 == nil) {
			id = api.FilterFirstAlbum(groupData.Artists, min-1, max+1) // Filtre les groupes en fonction
			groupData = IdToJson(mockData, id) // Recalcule groupData en fonction
		} else { // En cas d'erreur
			fmt.Println("error main firstAlbum")
			fmt.Println(err1, err2)
		}

		if len(location) > 0 { // Si un lieux de concert a été sélectionné
			id = api.FilterLocations(groupData.Locations, location) // Filtre les groupes en fonction
			groupData = IdToJson(mockData, id) // Recalcule groupData en fonction
		}

		id = []int{} // Réinitialise id
		membersFiltered := 0 // Signale si les membres ont été filtrés
		// Filtre les nombres de membres un à un
		if oneMember == "on" {
			id = append(id, api.FilterMembers(groupData.Artists, 1)...)
			membersFiltered++
		}
		if twoMembers == "on" {
			id = append(id, api.FilterMembers(groupData.Artists, 2)...)
			membersFiltered++
		}
		if threeMembers == "on" {// Recalcule groupData en fonction
			membersFiltered++
		}
		if fiveMembers == "on" {
			id = append(id, api.FilterMembers(groupData.Artists, 5)...)
			membersFiltered++
		}
		if sixMembers == "on" {
			id = append(id, api.FilterMembers(groupData.Artists, 6)...)
			membersFiltered++
		}
		if sevenMembers == "on" {
			id = append(id, api.FilterMembers(groupData.Artists, 7)...)
			membersFiltered++
		}
		
		if membersFiltered > 0 { // Si les membres ont été filtrés
			groupData = IdToJson(mockData, id) // Recalcule groupData en fonction
		}

		fmt.Println(id) // Affiche les ids des groupes filtrés dans le terminal
	}

	// Si une recherche a été saisie
	if len(input) > 0 {
		result := api.Search(input, groupData.Artists, groupData.Locations) // Recherche la saisie dans les données
		groupData = IdToJson(groupData, result) // Recalcule groupData en fonction
	}

	tmpl.Execute(w, groupData) // Execute le code html en fonction des changements de variables
}

// Récupère les données de l'api et les stockent dans une structure datasJson sous forme de json
func InitializeData() datasJson {
	var apiData datasJson
	json.Unmarshal(api.ExtractRawData(0), &apiData.Artists)
	json.Unmarshal(api.ExtractRawData(1), &apiData.Locations)
	json.Unmarshal(api.ExtractRawData(2), &apiData.Dates)
	apiData.Relations = api.GetConcerts(string(api.ExtractRawData(3)))
	apiData.Maps = maps

	return apiData
}

// Crée un datasJson contenant uniquement les données des groupes d'ids donnés
func IdToJson(groupData datasJson, id []int) datasJson {
	var data datasJson
	for _, i := range id {
		data.Artists = append(data.Artists, groupData.Artists[i])
		data.Locations.Index = append(data.Locations.Index, groupData.Locations.Index[i])
		data.Dates.Index = append(data.Dates.Index, groupData.Dates.Index[i])
		data.Relations = append(data.Relations, groupData.Relations[i])
		data.Maps = append(data.Maps, groupData.Maps[i])
	}

	return data
}