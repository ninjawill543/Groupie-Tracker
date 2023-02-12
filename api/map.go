//Ce fichier sert à générer une URL qui pointe vers une carte mondiale, qui contient des markers pour chaque lieu de concert d'un artiste

package apiFunctions

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

//Contient les informations de l'api GeoCoding de mapbox
type Response struct {
	Type     string   `json:"type"`
	Query    []string `json:"query"`
	Features []struct {
		ID         string   `json:"id"`
		Type       string   `json:"type"`
		PlaceType  []string `json:"place_type"`
		Relevance  int      `json:"relevance"`
		Properties struct {
			ShortCode string `json:"short_code"`
			Wikidata  string `json:"wikidata"`
		} `json:"properties"`
		Text              string    `json:"text"`
		PlaceName         string    `json:"place_name"`
		MatchingPlaceName string    `json:"matching_place_name"`
		Bbox              []float64 `json:"bbox"`
		Center            []float64 `json:"center"`
		Geometry          struct {
			Type        string    `json:"type"`
			Coordinates []float64 `json:"coordinates"`
		} `json:"geometry"`
		Context []struct {
			ID        string `json:"id"`
			ShortCode string `json:"short_code"`
			Wikidata  string `json:"wikidata"`
			Text      string `json:"text"`
		} `json:"context"`
	} `json:"features"`
	Attribution string `json:"attribution"`
}

// Cette fonction GetCoords reçoit un lieu de concert, qui est ensuite converti en coordonnées grâce à l'api Geocoding de MapBox, et renvoyé dans la fonction SendMarker

func GetCoords(place string) string {
    resp, err := http.Get("https://api.mapbox.com/geocoding/v5/mapbox.places/" + place + ".json?proximity=ip&limit=1&access_token=pk.eyJ1IjoibmluamF3aWxsNTQzIiwiYSI6ImNsZDRidWdrNjBvbDczcW9jajU5c3UxdXAifQ._jM-ztlL3-V9_0WjlFB01A")
    if err != nil {
        fmt.Println("No response from request")
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body) 

    var result Response
    if err := json.Unmarshal(body, &result); err != nil { 
        fmt.Println("Can not unmarshal JSON")
    }
	send := ""
    for _, rec := range result.Features {
		uno := fmt.Sprint(rec.Center[:1])
		dos := fmt.Sprint(rec.Center[1:])
		lat := uno[1 : len(uno)-1]
		long := dos[1 : len(dos)-1]
		send = (lat + "," + long)
    }

	return send 
}

// La fonction SendMarker prend une liste de lieux de concerts pour un artiste, qu'il renvoie une par une dans la fonction GetCoords pour convertir les lieux en coordonnés
// Ensuite, la liste de coordonnées est placé dans une URL afin de créer une carte avec pour marker chaque lieu de concert.
// Pour créer la carte, on utilise l'api Static Images de Mapbox

func SendMarker(places []string)string{
	markerbase := "pin-l-music+dd380f("
	endmarker := ""
	for _, element := range places {
		funccoord := GetCoords(element)
		endmarker = endmarker + markerbase + funccoord + ")," 
	}

	return "https://api.mapbox.com/styles/v1/ninjawill543/cld4bwm4c000h01tbd9cr3hun/static/" + string([]rune(endmarker)[:len(endmarker)-1]) + "/20,0,1/1000x1000?access_token=pk.eyJ1IjoibmluamF3aWxsNTQzIiwiYSI6ImNsZDRidWdrNjBvbDczcW9jajU5c3UxdXAifQ._jM-ztlL3-V9_0WjlFB01A"
}

func SendMaps(locationData Locations) []string {
	var urls []string
	for _, i := range locationData.Index {
		urls = append(urls, SendMarker(i.Locations))
	}

	return urls
}	