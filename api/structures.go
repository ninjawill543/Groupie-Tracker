/*
Contient les structures json utilisées par l'api donné : https://groupietrackers.herokuapp.com/api
*/

package apiFunctions

// Contient les informations générales des artistes/ groupe du lien : https://groupietrackers.herokuapp.com/api/artists
type Artists []struct {	
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

// Contient les lieux des concerts du lien : https://groupietrackers.herokuapp.com/api/locations
type Locations struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

// Contient les dates des concerts du lien : https://groupietrackers.herokuapp.com/api/dates
type Dates struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}