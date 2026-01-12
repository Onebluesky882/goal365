package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PlayerLoginLogRequest struct {
	PlayerID  uuid.UUID `bun:"player_id,type:uuid,notnull"`
	IPAddress string    `bun:"ip_address,notnull"`
	UserAgent string    `bun:"user_agent"`
}

type PlayerLoginLog struct {
	bun.BaseModel `bun:"table:player_login_logs"`

	ID        int64     `bun:"id,pk,autoincrement"`
	PlayerID  uuid.UUID `bun:"player_id,type:uuid,notnull"`
	IPAddress string    `bun:"ip_address,notnull"`
	UserAgent string    `bun:"user_agent"`
	CreatedAt time.Time `bun:"created_at,default:now()"`

	Player *Player `bun:"rel:belongs-to,join:player_id=id"`
}
