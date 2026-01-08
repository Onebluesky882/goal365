package lib

import (
	"mytipster/models"
)

func FilterBookmaker(items []models.Bookmaker, bookMarket string) []models.Bookmaker {
	var result []models.Bookmaker

	for _, bm := range items {
		if bm.Name == bookMarket {
			result = append(result, bm)
		}
	}
	return result
}
