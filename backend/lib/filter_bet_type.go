package lib

import (
	m "mytipster/models"
	"strings"
)

/*
Match Winner
Home/Away
Second Half Winner
Asian Handicap
Goals Over/Under
Goals Over/Under First Half
HT/FT Double
Exact Score
Highest Scoring Half
Correct Score - First Half
Double Chance
First Half Winner
Asian Handicap First Half
Double Chance - First Half
Odd/Even
Draw No Bet (1st Half)
Draw No Bet (2nd Half)
*/
func FilterBetType(
	bookmaker []m.Bookmaker,
	betsType []string,
) []m.Bet {

	var result []m.Bet
	betsTypeMap := make(map[string]struct{})

	for _, t := range betsType {
		betsTypeMap[strings.ToLower(t)] = struct{}{}
	}
	for _, bm := range bookmaker {
		for _, bet := range bm.Bets {
			if _, ok := betsTypeMap[strings.ToLower(bet.Name)]; ok {
				result = append(result, bet)
			}
		}
	}
	return result
}
