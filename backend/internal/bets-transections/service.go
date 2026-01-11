package betstransections

import (
	"context"
	"errors"
	"mytipster/lib/common"
	"mytipster/models"

	"github.com/uptrace/bun"
)

type TransactionService struct {
	db *bun.DB
}

func NewTransaction(db *bun.DB) *TransactionService {
	return &TransactionService{
		db: db,
	}
}

var req models.CreateTransactionRequest

func (s *TransactionService) InsertTransaction(ctx context.Context, body models.CreateTransactionRequest) (*models.Transaction, error) {
	playerID := body.PlayerId
	bets := body.Bets
	if len(bets) == 0 {
		return nil, errors.New("bets is empty")
	}

	// ---- calculate total ----
	var total int64
	for b := range bets {
		if bets[b].ID == 0 {
			bets[b].ID = common.Random10Digit()
		}
		total += bets[b].Amount
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	// ---- create transaction ----

	transaction := &models.Transaction{
		PlayerId: playerID,
		Type:     "BET",
		Total:    total,
		Status:   "PENDING",
	}

	if _, err := tx.NewInsert().
		Model(transaction).
		Exec(ctx); err != nil {
		return nil, err
	}

	// ---- attach transaction_id to bets ----
	for i := range bets {
		bets[i].TransactionID = transaction.ID
		bets[i].Status = "PENDING"
	}

	// ----- insert bets ------

	if _, err := tx.NewInsert().
		Model(&bets).Exec(ctx); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	transaction.Bets = bets

	return transaction, nil
}
