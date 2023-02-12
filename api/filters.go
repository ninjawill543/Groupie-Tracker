/*
Filtre les données
*/

package apiFunctions

import(
	"strconv"
	"fmt"
)

/* Filtre les doonnées en fonction de la date de création du groupe : retourne 
les groupes contenu dans artistData dont la création de groupe est entre les 
dates min et max données */
func FilterCreationDate(artistsData Artists, min int, max int) []int {
	var id []int
	for _, i := range artistsData { // Parcours les groupes donnés
		if (i.CreationDate > min) && (i.CreationDate < max) { // Vérifie la date de création du groupe
			id = append(id, i.ID-1)
		}
	}

	return id // Retourne un tableau contenant les ids des groupes filtrés
}

/* Filtre les doonnées en fonction de la date du premier album : retourne 
les groupes contenu dans artistData dont le premier album est sortie entre 
les dates min et max données */
func FilterFirstAlbum(artistsData Artists, min int, max int) []int {
	var id []int
	for _, i := range artistsData { // Parcours les groupes donnés
		// Convertie la date du premier album de string à int, et récupère uniquement l'année
		date, err := strconv.Atoi(i.FirstAlbum[len(i.FirstAlbum)-4:])  
		if err == nil {
			if (date > min) && (date < max) { // Vérifie la date du premier album
				id = append(id, i.ID-1)
			}
		} else { // En cas d'erreur
			fmt.Println("error FilterFirstAlbum")
		}
	}

	return id // Retourne un tableau contenant les ids des groupes filtrés
}

/* Filtre les doonnées en fonction du nombre de membres dans le groupe : retourne
les groupes ayant autant de membres que le nombres members donné */
func FilterMembers(artistsData Artists, members int) []int {
	var id []int
	for _, i := range artistsData { // Parcours les groupes donnés
		if len(i.Members) == members { // Vérifie le nombre de membres
			id = append(id, i.ID-1)
		}
	}

	return id // Retourne un tableau contenant les ids des groupes filtrés
}

/* Filtre les doonnées en fonction du nombre des lieux de concerts : retourne 
les groupes dont les données sont contenu dans locationsData dont les concerts 
ont eu lieux dans les positions contenues dans location */
func FilterLocations(locationsData Locations, location string) []int {
	var id []int
	for _, i := range locationsData.Index { // Parcours les données
		for _, k := range i.Locations { // Parcours les lieux données
			if k == location { // Vérifie si il y égalité
				id = append(id, i.ID-1)
				break
			}
		}
	}

	return id // Retourne un tableau contenant les ids des groupes filtrés
}