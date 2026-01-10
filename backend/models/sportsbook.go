package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type SportsBook struct {
	bun.BaseModel `bun:"table:sportsbooks"`

	Id        uuid.UUID `bun:"id,type:uuid,default:gen_random_uuid()"`
	MatchDate time.Time `bun:"match_date,pk,notnull"`

	FxId    int    `bun:"fx_id,notnull"`
	Home    string `bun:"home"`
	Away    string `bun:"away"`
	League  string `bun:"league"`
	Country string `bun:"country"`

	Bet []Bet `bun:"bet,type:jsonb"`

	CreatedAt time.Time `bun:"created_at,default:now()"`
	UpdatedAt time.Time `bun:"updated_at,nullzero"`
}
