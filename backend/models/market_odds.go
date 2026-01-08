package models

import "github.com/google/uuid"

type MarketOdds struct {
	Id        uuid.UUID        `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	FxId      int              `bun:"fx_id,notnull" json:"fx_id"`
	Home      string           `bun:"-" json:"home"`      // ignore Bun, json ใช้งานได้
	Away      string           `bun:"-" json:"away"`      // ignore Bun
	League    string           `bun:"-" json:"league"`    // ignore Bun
	Country   string           `bun:"-" json:"country"`   // ignore Bun
	Bookmaker []Bookmaker      `bun:"-" json:"bookmaker"` // ignore Bun, append API data
	Bet       []Bet            `bun:"-" json:"bet"`
	Response  *FixtureResponse `bun:"-" json:"response,omitempty"`
}
