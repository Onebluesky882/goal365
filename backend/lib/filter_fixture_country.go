package lib

import m "mytipster/models"

func FilterCountry(fx *m.FixtureResponse, countries []string) bool {
	if fx == nil {
		return false
	}
	for _, c := range countries {
		if fx.League.Country == c {
			return true
		}
	}
	return false
}
