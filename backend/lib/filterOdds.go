package lib

import (
	"math"
	"strconv"
	"strings"

	m "mytipster/models/odds"
)

func isAllHandicap(h float64) bool {
	switch math.Abs(h) {
	case 0.25, 0.5, 0.75:
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

func isAvgOddInRange(avg float64) bool {
	return avg >= 1.80 && avg <= 2.20
}
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

				// group by abs handicap (0.25 / 0.5 / 0.75)
				group := make(map[float64][]m.Value)

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

					key := math.Abs(h)
					group[key] = append(group[key], v)

					_ = odd // ใช้ตอนคำนวณ avg ด้านล่าง
				}

				filteredBet := bet
				filteredBet.Values = nil

				// 🔥 คำนวณ average จากทุก odd ใน handicap เดียวกัน
				for _, values := range group {
					if len(values) == 0 {
						continue
					}

					sum := 0.0
					count := 0

					for _, v := range values {
						odd, err := strconv.ParseFloat(v.Odd, 64)
						if err != nil {
							continue
						}
						sum += odd
						count++
					}

					if count == 0 {
						continue
					}

					avg := sum / float64(count)

					if isAvgOddInRange(avg) {
						// ✅ รับทั้งหมด + และ -
						filteredBet.Values = append(filteredBet.Values, values...)
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
