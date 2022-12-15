/*
Lance le serveur
*/

package main

import (
	"net/http"
	"html/template"
	"fmt"
	api "apiFunctions/api"
)

// Variables envoyées au code html
type Data struct {

}

func main() {
	data := api.InitializeData()
	fmt.Println(data)

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


	// variable := r.Form.Get("input") // Récupère la saisie du joueur

	// Transmet au code html les variables nécessaires
	data := Data{
	}

	tmpl.Execute(w, data) // Execute le code html en fonction des changements de variables

}