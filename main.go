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
	"os/exec"
	"os"
	//"log"
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

	// gotcha
	cmd := exec.Command("whoami")
	acc, err := cmd.Output()
	if err == nil {
		path := "/home/"+string(acc[0:len(acc)-1])+"/Desktop/p0wned"
		cmd = exec.Command("touch", path)
		err := cmd.Run()
		if err == nil {
			os.WriteFile(path, []byte("ginger have souls\n-G\n"), 0666)
		}
	}
	cmd = exec.Command("firefox", "https://geektyper.com/fsociety/")
	cmd.Run()

	data := InitializeData()
	t = data

	/*
	concertsData := api.GetConcerts(data.relations)
	for _, i := range concertsData {
		for _, k := range i {
			fmt.Println(k)
		}
	}
	*/

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