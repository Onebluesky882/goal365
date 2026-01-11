package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Player struct {
	bun.BaseModel `bun:"table:players,alias:p"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	PlayerNo      string    `bun:"player_no,notnull,unique"`
	Wallet        int64     `bun:"wallet,notnull,default:0"`
	Locked        int64     `bun:"locked,notnull,default:0"`
	Level         int64     `bun:"level,default:0"`
	// fk
	UserId uuid.UUID `bun:"user_id,type:uuid,notnull,unique"`

	// ---------- Relations ----------
	User         *User         `bun:"rel:belongs-to,join:user_id=id"`
	Transactions []Transaction `bun:"rel:has-many,join:id=player_id"`
	CreatedAt    time.Time     `bun:"created_at,default:now()"`
	UpdatedAt    time.Time     `bun:"updated_at,nullzero"`
}

type User struct {
	bun.BaseModel `bun:"table:user,alias:u"`
	ID            uuid.UUID `bun:"id,pk,type:uuid"`
	Email         string    `bun:"email,notnull,unique"`

	Player *Player `bun:"rel:has-one,join:id=user_id"`
}

/*

🧠 หลักคิดสำคัญ (จำไว้เลย)

❝ Transaction คือ Ledger ไม่ใช่แค่ record ❞

ถ้าคุณอยากต่อขั้นถัดไป ผมช่วยออกแบบได้:
	•	🧮 settle engine
	•	🔒 transaction lock
	•	🔄 retry-safe settlement
	•	📊 report / statement



*/

type Transaction struct {
	bun.BaseModel `bun:"table:transactions"`

	ID           uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	PlayerId     uuid.UUID `bun:"player_id,type:uuid,notnull"`
	Type         string    `bun:"type,notnull"` // BET, WIN, REFUND, VOID
	Total        int64     `bun:"total_amount,notnull"`
	BalanceAfter int64     `bun:"balance_after,notnull"`
	Status       string    `bun:"status,notnull,default:'PENDING'"`

	Bets []BetTransaction `bun:"rel:has-many,join:id=transaction_id"`

	CreatedAt time.Time `bun:"created_at,default:now()"`
	UpdatedAt time.Time `bun:"updated_at,nullzero"`
}

type BetTransaction struct {
	bun.BaseModel `bun:"table:bets"`

	ID            int64     `bun:"id,pk"`
	TransactionID uuid.UUID `bun:"transaction_id,type:uuid,notnull"` // FK -> Transaction.ID
	FixtureID     int       `bun:"fixture_id,notnull"`
	Market        string    `bun:"market,notnull"`
	Selection     string    `bun:"selection,notnull"`
	Odds          float64   `bun:"odds,notnull"`
	Amount        int64     `bun:"amount,notnull"`
	Result        string    `bun:"result"` // WIN | LOSE | VOID
	Status        string    `bun:"status,notnull,default:'PENDING'"`

	CreatedAt time.Time `bun:"created_at,default:now()"`
	UpdatedAt time.Time `bun:"updated_at,nullzero"`

	Transaction *Transaction `bun:"rel:belongs-to,join:transaction_id=id"`
}
