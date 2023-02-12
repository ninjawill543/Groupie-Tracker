/*
Fonctions de recherches dans les données reçut par l'api en fonction de saisies de l'utilisateur
*/

package apiFunctions

import (
	"fmt"
	"strconv"
)

/* Recherche une donnée égale la saisie de l'utilisateur parmis les données suivantes :
nom du groupe, année du premier album, année de création du groupe, membres du groupe, lieux de concerts
Recherche aussi une donnée similaire à la saisir de l'utilisateur parmis les noms de groupes.
*/
func Search(input string, artistsData Artists, artistsLocations Locations) []int {
	input = Minimalize(input) // Transforme les majuscule de la saisie de l'utilsateur en minuscule
	var similarities []float64 // Contient les scores de similarités (cf description de CheckSimilarities)
	var id []int // Contient les ids des groupes avec donnée égale à une saisie
	for _, i := range artistsData { // Parcours les données
		// Si la saisie est égale au nom de groupe, l'année du premier album ou l'année de création du groupe
		if (input == Minimalize(i.Name)) || (input == i.FirstAlbum[len(i.FirstAlbum)-4:]) || (input == strconv.Itoa(i.CreationDate)) {
			id = append(id, i.ID-1) // Ajoute l'id
		} else { // Sinon
			for _, k := range i.Members { // Parcours les membres du groupe
				if input == Minimalize(k) { // Si la saisie est égale au nom d'un membre du groupe
					id = append(id, i.ID-1) // Ajoute l'id
					break
				}
			}
			if (len(id) == 0) || (id[len(id)-1] != i.ID-1) { // Si l'id du groupe considéré n'est pas encore ajouté
				for _, k := range artistsLocations.Index[i.ID-1].Locations { // Parcours les lieux de concerts
					if input == Minimalize(k) { // Si un lieux est égal à la saisie
						id = append(id, i.ID-1) // Ajoute l'id
						break
					}
				}
			}
		}
		similarities = append(similarities, CheckSimilarities(input, Minimalize(i.Name))) // Ajoute le score de similarité du groupe considéré
	}

	fmt.Println(id) // Affiche les solutions exactes dans le terminal

	var max float64 = 0 // Contient le score de similarité max actuel
	maxId := 0 // Contient l'id associé au score de similarité max actuel
	var maxSimilarities [12]int // Tableau contenant les 12 groupes avec le score de similarité le plus élevé
	for i := range maxSimilarities { // Initialise maxSimilarities
		maxSimilarities[i] = -1
	}
	index := 0 // Contient l'index actuel de maxSimilarities
	for index < len(maxSimilarities) { // Tant que l'index est inférieur à la longueur de maxSimilarities
		for i1, i2 := range similarities { // Parcous similarities
			if i2 > max { // Si le score de simlilarité actuel est supérieur au max
				// Parcours les groupes aux scores de similarités les plus élévés trouvé précédemments
				for k1, k2 := range maxSimilarities {
					if i1 == k2 { // Si le groupe est déjà dans maxSimilarities, ignore le groupe
						break
					} else if k1 == len(maxSimilarities)-1 { // Sinon, devient le nouveau max
						maxId = i1
						max = similarities[maxId]
					}
				}
			}
		}
		if max > 0 { // Si le score de similarité du max est supérieur à 0
			maxSimilarities[index] = maxId // L'ajoute
		} else { // Sinon (plus de groupe avec score de similarité au-dessus de 0)
			break // Interrompt la boucle
		}
		index++ // Une valeur de plus dans maxSimilarities
		max = 0 // Réinitialise le max
	}
	
	var solution []int // Contient la solution a renvoyer
	for _, i := range maxSimilarities { // Parcours les max
		if (i > -1) && (similarities[i] > 0) { // Si le max actuel existe
			solution = append(solution, i) // L'ajoute à la solution
		}
	}
	if len(id) > 0 { // Si des solutions exactes ont été trouvé
		for _, i := range id { // Parcours les solutions exactes
			for k1, k2 := range solution { // Parcous les solutions à renvoyer
				if i == k2 { // Vérifie la solution exacte est déjà contenue
					// La met à l'avant des solutions à renvoyer (et supprime l'autre itération)
					solution = append(solution[0:k1], solution[k1+1:]...) 
					break
				}
			}
			solution = append([]int{i}, solution...) // Sinon, l'ajoute à l'avant
		}
	}

	fmt.Println(solution) // Affiche la solution dans le terminal

	return solution // renvoie les solutions
}

// Transforme les majuscule en minuscule d'un string donné
func Minimalize(s string) string {
	var solu string
	for i, _ := range s {
		if (s[i] >= 'A') && (s[i] <= 'Z') {
			solu += string(s[i]+('a'-'A'))
		} else {
			solu += string(s[i])
		}
	}

	return solu
}

/* Retourne le score de similarité entre deux strings donnés :
un point par charactère en commun, divisé par la longueur du second string */
func CheckSimilarities(s1 string, s2 string) float64 {
	similarities := 0
	for _, k := range s1 {
		for _, j := range s2 {
			if k == j {
				similarities ++
				break
			}
		}
	}

	return float64(similarities)/float64(len(s2))
}