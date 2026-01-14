package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

/*

🧠 หลักคิดสำคัญ (จำไว้เลย)

❝ Transaction คือ Ledger ไม่ใช่แค่ record ❞

ถ้าคุณอยากต่อขั้นถัดไป ผมช่วยออกแบบได้:
	•	🧮 settle engine
	•	🔒 transaction lock
	•	🔄 retry-safe settlement
	•	📊 report / statement

2. Transaction Boundary

คุณเริ่มถาม BeginTx แล้ว 👍
ขั้นต่อไปคือ:
	•	atomic insert (transaction + bets)
	•	rollback on error
	•	balance consistency
*/

type Player struct {
	bun.BaseModel `bun:"table:players,alias:p"`

	ID           uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	PlayerNo     int64     `bun:"player_no,notnull,unique" json:"playerNo"`
	Level        int8      `bun:"level,default:1" json:"level"`
	Name         string    `bun:"name" json:"name"`
	Bio          string    `bun:"bio" json:"bio"`
	ImageUrl     string    `bun:"image_url" json:"imageUrl"`
	Wallet       int64     `bun:"wallet,notnull,default:0" json:"wallet"`
	WalletLocked int64     `bun:"wallet_locked" json:"walletLocked"`
	Locked       bool      `bun:"locked,default:false" json:"locked"`
	UserId       string    `bun:"user_id,notnull" json:"userId"`

	User         *User         `bun:"rel:belongs-to,join:user_id=id" json:"-"`
	Transactions []Transaction `bun:"rel:has-many,join:player_no=player_no" json:"transactions,omitempty"`

	CreatedAt time.Time `bun:"created_at,default:now()" json:"createdAt"`
	UpdatedAt time.Time `bun:"updated_at,nullzero" json:"updatedAt"`
}

type Transaction struct {
	bun.BaseModel `bun:"table:transactions"`

	Id       uuid.UUID `bun:"transaction_id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	PlayerNo int64     `bun:"player_no,notnull" json:"playerNo"`

	Type    string `bun:"type,notnull" json:"type"`
	Total   int64  `bun:"total_amount,notnull" json:"total"`
	Status  string `bun:"status,notnull,default:'PENDING'" json:"status"`
	Settled bool   `bun:"settled,default:false"`

	Bets []BetTransaction `bun:"rel:has-many,join:transaction_id=transaction_id" json:"bets,omitempty"`

	Player *Player `bun:"rel:belongs-to,join:player_no=player_no" json:"player,omitempty"`

	CreatedAt time.Time `bun:"created_at,default:now()" json:"createdAt"`
	UpdatedAt time.Time `bun:"updated_at,nullzero" json:"updatedAt,omitempty"`
}

type BetTransaction struct {
	bun.BaseModel `bun:"table:bets"`

	BetId         int64     `bun:"bet_id,pk" json:"betId"`
	TransactionID uuid.UUID `bun:"transaction_id,type:uuid,notnull" json:"transactionId"`
	FixtureID     int       `bun:"fixture_id,notnull" json:"fixtureId"`
	Market        string    `bun:"market,notnull" json:"market"`
	Selection     string    `bun:"selection,notnull" json:"selection"`
	Odds          float64   `bun:"odds,notnull" json:"odds"`
	Amount        int64     `bun:"amount,notnull" json:"amount"`

	Result  string `bun:"result" json:"result"`
	Status  string `bun:"status,notnull,default:'PENDING'" json:"status"`
	Settled bool   `bun:"settled,default:false" json:"settled"`

	Transaction *Transaction `bun:"rel:belongs-to,join:transaction_id=transaction_id" json:"transaction,omitempty"`

	CreatedAt time.Time `bun:"created_at,default:now()" json:"createdAt"`
	UpdatedAt time.Time `bun:"updated_at,nullzero" json:"updatedAt,omitempty"`
}
type UpdateTransactionRequest struct {
	PlayerNo int64 `json:"player_no"`

	Bets []BetTransaction `json:"bets"`
}

type CreateTransactionRequest struct {
	PlayerNo int64              `json:"player_no"`
	Type     string             `json:"type"`
	Bets     []CreateBetRequest `json:"bets"`
}

type CreateBetRequest struct {
	FixtureID int     `json:"fixture_id"`
	Market    string  `json:"market"`
	Selection string  `json:"selection"`
	Odds      float64 `json:"odds"`
	Amount    int64   `json:"amount"`
}

type CreatePlayerRequest struct {
	Name     string `bun:"name" json:"name"`
	Bio      string `bun:"bio" json:"bio"`
	PlayerNo int64  `bun:"player_no,notnull,unique"`
	Wallet   int64  `bun:"wallet,notnull,default:0"`
	UserID   string `bun:"user_id" json:"user_id"`
}

type UpdatePlayerRequest struct {
	Name string `json:"name" validate:"required"`
}

type SetWalletRequest struct {
	PlayerId uuid.UUID `json:"player_id" validate:"required"`
	Amount   int64     `json:"amount" validate:"required"`
}
