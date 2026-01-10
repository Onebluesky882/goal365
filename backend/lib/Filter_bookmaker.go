package lib

import (
	"mytipster/models"
	"strings"
)

/*
{id: 7, name: 'William Hill', bets: Array(14)}
 {id: 8, name: 'Bet365', bets: Array(17)}
 {id: 2, name: 'Marathonbet', bets: Array(14)}
 {id: 16, name: 'Unibet', bets: Array(18)}
 {id: 4, name: 'Pinnacle', bets: Array(7)}
 {id: 5, name: 'SBO', bets: Array(9)}
 {id: 11, name: '1xBet', bets: Array(30)}
 {id: 32, name: 'Betano', bets: Array(17)}
 {id: 34, name: 'Superbet', bets: Array(9)}
 {id: 9, name: 'Dafabet', bets: Array(1)}


*/

func FilterBookmaker(items []models.Bookmaker, bookMarket []string) []models.Bookmaker {
	var result []models.Bookmaker

	for _, bm := range items {
		name := strings.TrimSpace(strings.ToLower(bm.Name))
		// func ContainsString(arr []string, target string) bool

		for _, target := range bookMarket {
			if name == strings.ToLower(target) {
				result = append(result, bm)
				break
			}
		}
	}
	return result
}
