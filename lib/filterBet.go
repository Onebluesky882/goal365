package lib

import (
	m "mytipster/models/odds"
)
func FilterBookMarket(root *m.RootOdds, bookmakerName string) map[int][]m.Bet {
	result := make(map[int][]m.Bet)

	for _, resp := range root.Response {
		for _, bm := range resp.Bookmakers {
			if bm.Name != bookmakerName {
				continue // skip if not the target bookmaker
			}

			for _, bet := range bm.Bets {
				if bet.Name == "Asian Handicap" {
					result[resp.Fixture.ID] = []m.Bet{bet} // add directly
					break // ไม่ต้อง loop bet ที่เหลือ เพราะเรากรองเฉพาะ Asian Handicap
				}
			}
			break // ไม่ต้อง loop bookmaker ตัวอื่นสำหรับ fixture นี้
		}
	}

	return result
}