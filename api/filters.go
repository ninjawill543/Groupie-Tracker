package apiFunctions

import(
	"strconv"
	"fmt"
)

func FilterCreationDate(artistsData Artists, min int, max int) []int {
	var id []int
	for _, i := range artistsData {
		if (i.CreationDate > min) && (i.CreationDate < max) {
			id = append(id, i.ID-1)
		}
	}

	return id
}

func FilterFirstAlbum(artistsData Artists, min int, max int) []int {
	var id []int
	for _, i := range artistsData {
		date, err := strconv.Atoi(i.FirstAlbum[len(i.FirstAlbum)-4:])
		if err == nil {
			if (date > min) && (date < max) {
				id = append(id, i.ID-1)
			}
		} else {
			fmt.Println("error FilterFirstAlbum")
		}
	}

	return id
}

func FilterMembers(artistsData Artists, members int) []int {
	var id []int
	for _, i := range artistsData {
		if len(i.Members) == members {
			id = append(id, i.ID-1)
		}
	}

	return id
}

func FilterLocations(locationsData Locations, location string) []int {
	var id []int
	for _, i := range locationsData.Index {
		for _, k := range i.Locations {
			if k == location {
				id = append(id, i.ID)
				break
			}
		}
	}

	return id
}