package apiFunctions

import (
)

func Search(input string, artistsData Artists) []int {
	input = Minimalize(input)
	var similarities []int
	id := -1
	for _, i := range artistsData {
		name := Minimalize(i.Name)
		if input == Minimalize(name) {
			id = i.ID
		}
		similarities = append(similarities, CheckSimilarities(input, name))
	}

	max := 0
	maxId := 0
	var maxSimilarities [5]int
	for i := range maxSimilarities {
		maxSimilarities[i] = -1
	}
	index := 0
	for index < 5 {
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
		maxSimilarities[index] = maxId
		index++
		max = 0
	}

	var solution []int
	for _, i := range maxSimilarities {
		solution = append(solution, i)
	}
	if id > -1 {
		for i1, i2 := range solution {
			if (i2 == id) && (i1 != 0) {
				solution = append(solution[0:i1], solution[i1+1:]...)
				solution = append([]int{id}, solution...)
			}
		}
	}

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