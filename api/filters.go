package apiFunctions

import(
	
)

func FilterCreationDate(artistsData Artists, mode int, date int) []int {
	var id []int
	for _, i := range artistsData {
		switch mode {
		case 0:
			if i.CreationDate == date {
				id = append(id, i.ID)
			}
		case 1:
			if i.CreationDate < date {
				id = append(id, i.ID)
			}
		case 2:
			if i.CreationDate > date {
				id = append(id, i.ID)
			}
		}
	}

	return id
}

func FilterFirstAlbum(artistsData Artists, mode int, date string) []int {
	var id []int
	for _, i := range artistsData {
		if i.FirstAlbum == date {
			id = append(id, i.ID)
		}
	}

	return id
}

func FilterConcertDate(datesData Dates, mode int, date string) []int {
	var id []int
	for _, i := range datesData.Index {
		for _, k := range i.Dates {
			if k == date {
				id = append(id, i.ID)
			}
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