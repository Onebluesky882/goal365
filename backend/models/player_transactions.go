package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Player struct {
	bun.BaseModel `bun:"table:players,alias:p"`

	ID           uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	PlayerNo     int64     `bun:"player_no,notnull,unique,default:0" json:"playerNo"`
	Name         string    `bun:"name" json:"name"`
	Bio          string    `bun:"bio" json:"bio"`
	ImageUrl     string    `bun:"image_url" json:"imageUrl"`
	Wallet       int64     `bun:"wallet,notnull,default:0" json:"wallet"`
	Level        int64     `bun:"level,default:0" json:"level"`
	Exp          int64     `bun:"exp" json:"exp"`
	WalletLocked int64     `bun:"wallet_locked" json:"walletLocked"`  // เงินที่กันไว้
	Locked       bool      `bun:"locked,default:false" json:"locked"` // lock account

	UserId string `bun:"user_id,notnull" json:"userId"`

	// ---------- Relations ----------
	User         *User         `bun:"rel:belongs-to,join:user_id=id" json:"-"`
	Transactions []Transaction `bun:"rel:has-many,join:id=player_id" json:"transactions,omitempty"`

	CreatedAt time.Time `bun:"created_at,default:now()" json:"createdAt"`
	UpdatedAt time.Time `bun:"updated_at,nullzero" json:"updatedAt"`
}

type User struct {
	bun.BaseModel `bun:"table:user"`
	ID            string   `bun:"id,pk"`
	Email         string   `bun:"email,notnull,unique"`
	Players       []Player `bun:"rel:has-many,join:id=user_id"`
}

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

type Transaction struct {
	bun.BaseModel `bun:"table:transactions"`

	ID           uuid.UUID        `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	BillId       int64            `bun:"bill_id"`
	PlayerId     uuid.UUID        `bun:"player_id,type:uuid,notnull"`
	Type         string           `bun:"type,notnull"` // BET, WIN, REFUND, VOID
	Total        int64            `bun:"total_amount,notnull"`
	BalanceAfter int64            `bun:"balance_after,notnull"`
	Status       string           `bun:"status,notnull,default:'PENDING'"`
	MixParlay    bool             `bun:"mix_parlay,notnull,default:false"`
	Bets         []BetTransaction `bun:"rel:has-many,join:id=transaction_id"`
	Settled      bool             `bun:"settled,default:false"`
	CreatedAt    time.Time        `bun:"created_at,default:now()"`
	UpdatedAt    time.Time        `bun:"updated_at,nullzero"`
}

type BetTransaction struct {
	bun.BaseModel `bun:"table:bets"`

	ID            int64     `bun:"id,pk"`
	BetId         int64     `bun:"bet_id"`
	TransactionID uuid.UUID `bun:"transaction_id,type:uuid,notnull,on_delete:cascade "` // FK -> Transaction.ID
	FixtureID     int       `bun:"fixture_id,notnull"`
	Market        string    `bun:"market,notnull"`
	Selection     string    `bun:"selection,notnull"`
	Odds          float64   `bun:"odds,notnull"`
	Amount        int64     `bun:"amount,notnull"`
	Result        string    `bun:"result"` // WIN | LOSE | VOID
	Status        string    `bun:"status,notnull,default:'PENDING'"`
	Settled       bool      `bun:"settled,default:false"`
	CreatedAt     time.Time `bun:"created_at,default:now()"`
	UpdatedAt     time.Time `bun:"updated_at,nullzero"`

	Transaction *Transaction `bun:"rel:belongs-to,join:transaction_id=id"`
}

type UpdateTransactionRequest struct {
	BillId   int64     `json:"bill_id"`
	PlayerId uuid.UUID `json:"player_id"`

	Bets []BetTransaction `json:"bets"`
}

type CreateTransactionRequest struct {
	PlayerId uuid.UUID        `json:"player_id"`
	Bets     []BetTransaction `json:"bets"`
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
