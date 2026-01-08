package sportbook

import (
	"log"
	"mytipster/internal/fixtures"
	"mytipster/lib"
	"mytipster/models"
	"strconv"
)

type MarketOdds struct {
}

func GetMarketOdds(date string) ([]models.MarketOdds, error) {
	ids, err := fixtures.GetIds(date)
	if err != nil {
		return nil, err
	}
	var markets []models.MarketOdds

	for _, id := range ids {
		// query odds from API
		fxResp, err := fixtures.QueryFixtureId(strconv.Itoa(id))
		if err != nil {
			log.Printf("❌ QueryFixtureId failed for fixture %d: %v", id, err)
			continue // skip fixture นี้
		}

		if fxResp == nil {
			log.Printf("⚠️ No fixture data for id %d", id)
			continue
		}

		oddsResp, err := fixtures.QueryFixtureOdds(strconv.Itoa(id))

		if err != nil {
			log.Printf("❌ QueryFixtureOdds failed for fixture %d: %v", id, err)
			continue
		}

		if oddsResp == nil {
			log.Printf("⚠️ No odds data for fixture %d", id)
			continue
		}

		market := models.MarketOdds{
			FxId:     id,
			Response: fxResp,
			Country:  fxResp.League.Country,
			League:   fxResp.League.Name,
			Home:     fxResp.Teams.Home.Name,
			Away:     fxResp.Teams.Away.Name,
			Bet:      []models.Bet{},
		}

		filtered := lib.FilterBookmaker(oddsResp, "Betano")
		market.Bookmaker = filtered

		// Flatten bets จาก bookmakers ที่กรองแล้ว
		for _, bm := range filtered {
			market.Bet = append(market.Bet, bm.Bets...)
		}
		markets = append(markets, market)
		break
		// stop after first fixture

	}

	return markets, nil

}
