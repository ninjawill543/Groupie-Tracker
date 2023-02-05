package apiFunctions

import (
	"fmt"
	"strconv"
)

func Search(input string, artistsData Artists, artistsLocations Locations) []int {
	input = Minimalize(input)
	var similarities []int
	var id []int
	for _, i := range artistsData {
		if (input == Minimalize(i.Name)) || (input == i.FirstAlbum[len(i.FirstAlbum)-4:]) || (input == strconv.Itoa(i.CreationDate)) {
			id = append(id, i.ID-1)
		} else {
			for _, k := range i.Members {
				if input == Minimalize(k) {
					id = append(id, i.ID-1)
					break
				}
			}
			if (len(id) == 0) || (id[len(id)-1] != i.ID-1) {
				for _, k := range artistsLocations.Index[i.ID-1].Locations {
					if input == Minimalize(k) {
						id = append(id, i.ID-1)
						break
					}
				}
			}
		}
		similarities = append(similarities, CheckSimilarities(input, i.Name))
	}

	fmt.Println(id)

	max := 0
	maxId := 0
	var maxSimilarities [10]int
	for i := range maxSimilarities {
		maxSimilarities[i] = -1
	}
	index := 0
	for index < len(maxSimilarities) {
		for i1, i2 := range similarities {
			if i2 > max {
				for k1, k2 := range maxSimilarities {
					if i1 == k2 {
						break
					} else if k1 == len(maxSimilarities)-1 {
						maxId = i1
						max = similarities[maxId]
					}
				}
			}
		}
		if similarities[maxId] > 0 {
			for _, i := range maxSimilarities {
				if maxId == i {
					break
				} else if i == maxSimilarities[len(maxSimilarities)-1] {
					maxSimilarities[index] = maxId
				}
			}
		}
		index++
		max = 0
	}

	fmt.Println(similarities)

	var solution []int
	for _, i := range maxSimilarities {
		if (i > -1) && (similarities[i] > 0) {
			solution = append(solution, i)
		}
	}
	if len(id) > 0 {
		for _, i := range id {
			for k1, k2 := range solution {
				if i == k2 {
					solution = append(solution[0:k1], solution[k1+1:]...)
					break
				}
			}
			solution = append([]int{i}, solution...)
		}
	}

	fmt.Println(solution)

	return solution
}

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

func CheckSimilarities(s1 string, s2 string) int {
	similarities := 0
	for _, k := range s1 {
		for _, j := range s2 {
			if k == j {
				similarities ++
				break
			}
		}
	}

	return similarities
}

func Minimum(s [5]int) int {
	max := s[0]
	for _, i := range s {
		if i > max {
			max = i
		}
	}

	return max
}