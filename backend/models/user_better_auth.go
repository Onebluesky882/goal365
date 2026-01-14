package models

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:user"`
	ID            string   `bun:"id,pk"`
	Email         string   `bun:"email,notnull,unique"`
	Players       []Player `bun:"rel:has-many,join:id=user_id"`
}
