package oddsmap

import (
	"math"
	"strconv"
	"strings"

	m "mytipster/models/odds"
)

func isAllHandicap(h float64) bool {
	switch math.Abs(h) {
	case 0.25, 0.5, 0.75, 1.0:
		return true
	default:
		return false
	}
}

func parseHandicap(val string) (float64, bool) {
	parts := strings.Fields(val)
	if len(parts) < 2 {
		return 0, false
	}

	h, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return 0, false
	}

	return h, true
}

func isOddInRange(odd float64) bool {
	return odd >= 1.75 && odd <= 2.20
}

// ฟังก์ชันใหม่สำหรับกรอง OddsMap
func FilterOddsMap(oddsMap m.OddsMap) map[int][]m.Bet {
	result := make(map[int][]m.Bet)

	// วนลูปแต่ละ fixture
	for fixtureIDStr, bets := range oddsMap {
		// แปลง fixture ID จาก string เป็น int
		fixtureID, err := strconv.Atoi(fixtureIDStr)
		if err != nil {
			continue
		}

		var filteredBets []m.Bet

		// วนลูปแต่ละ bet type
		for _, bet := range bets {
			// เช็คว่าเป็น Asian Handicap หรือไม่
			if strings.TrimSpace(bet.Name) != "Asian Handicap" {
				continue
			}

			filteredBet := bet
			filteredBet.Values = nil

			// วนลูปแต่ละ value
			for _, v := range bet.Values {
				valStr, ok := v.Value.(string)
				if !ok {
					continue
				}

				// Parse handicap value
				h, ok := parseHandicap(valStr)
				if !ok || !isAllHandicap(h) {
					continue
				}

				// Parse odd
				odd, err := strconv.ParseFloat(v.Odd, 64)
				if err != nil {
					continue
				}

				// 🔥 เงื่อนไขหลัก: odd ต้องอยู่ระหว่าง 1.80-2.20
				if isOddInRange(odd) {
					filteredBet.Values = append(filteredBet.Values, v)
				}
			}

			// ถ้ามี values ที่ผ่านเงื่อนไข ให้เพิ่มเข้า result
			if len(filteredBet.Values) > 0 {
				filteredBets = append(filteredBets, filteredBet)
			}
		}

		// ถ้ามี bets ที่ผ่านเงื่อนไข ให้เพิ่มเข้า result
		if len(filteredBets) > 0 {
			result[fixtureID] = filteredBets
		}
	}

	return result
}

// ฟังก์ชันเดิมสำหรับ RootOdds (เก็บไว้ backward compatible)
func FilterOdds(root *m.RootOdds, bookmakerName string) map[int][]m.Bet {
	result := make(map[int][]m.Bet)

	for _, resp := range root.Response {
		for _, bm := range resp.Bookmakers {
			if bm.Name != bookmakerName {
				continue
			}

			var filteredBets []m.Bet

			for _, bet := range bm.Bets {
				if strings.TrimSpace(bet.Name) != "Asian Handicap" {
					continue
				}

				filteredBet := bet
				filteredBet.Values = nil

				for _, v := range bet.Values {
					valStr, ok := v.Value.(string)
					if !ok {
						continue
					}

					h, ok := parseHandicap(valStr)
					if !ok || !isAllHandicap(h) {
						continue
					}

					odd, err := strconv.ParseFloat(v.Odd, 64)
					if err != nil {
						continue
					}

					if isOddInRange(odd) {
						filteredBet.Values = append(filteredBet.Values, v)
					}
				}

				if len(filteredBet.Values) > 0 {
					filteredBets = append(filteredBets, filteredBet)
				}
			}

			if len(filteredBets) > 0 {
				result[resp.Fixture.ID] = filteredBets
			}

			break
		}
	}

	return result
}
