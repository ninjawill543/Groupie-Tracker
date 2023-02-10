package map

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
	"html/template"
	//"os"
)

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

func SendMarker(places []string)string{
	//places := []string{}
	markerbase := "pin-l-music+dd380f("
	endmarker := ""
	for _, element := range places {
		funccoord := GetCoords(element)
		endmarker = endmarker + markerbase + funccoord + ")," 
	}
	return(string([]rune(endmarker)[:len(endmarker)-1]))
}

