package betstransections

import (
	"context"
	"errors"
	"mytipster/lib/common"
	"mytipster/models"

	"github.com/google/uuid"
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

func (s *TransactionService) getTransaction(ctx context.Context, req models.UpdateTransactionRequest) (*models.Transaction, error) {
	tx := new(models.Transaction)

	err := s.db.NewSelect().
		Model(tx).
		Relation("Bets", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.
				Order("created_at ASC")
		}).
		Where("player_no = ?", req.PlayerNo).
		Where("bill_id = ?", req.BillId).
		Scan(ctx)

	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (s *TransactionService) InsertTransaction(
	ctx context.Context,
	bets []models.BetTransaction,
	playerNo int64,
) (*models.Transaction, error) {

	// ---------- validate ----------
	if len(bets) == 0 {
		return nil, errors.New("bets is empty")
	}

	// ---------- calculate total ----------
	var total int64
	for i := range bets {
		if bets[i].Amount <= 0 {
			return nil, errors.New("invalid bet amount")
		}
		total += bets[i].Amount
	}

	// ---------- check player exists ----------
	exists, err := s.db.NewSelect().
		Model((*models.Player)(nil)).
		Where("player_no = ?", playerNo).
		Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("player not found")
	}

	// ---------- begin transaction ----------
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	// safety rollback
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	// ---------- HOLD WALLET ----------
	if err = s.holdWalletTx(ctx, tx, playerNo, total); err != nil {
		return nil, err
	}

	// ---------- create transaction ----------
	transaction := &models.Transaction{
		PlayerNo: playerNo,
		Type:     "BET",
		Total:    total,
		Status:   "PENDING",
		Settled:  false,
	}

	if _, err = tx.NewInsert().
		Model(transaction).
		Exec(ctx); err != nil {
		return nil, err
	}

	// ---------- attach transaction_id to bets ----------
	for i := range bets {
		bets[i].TransactionID = transaction.Id
		bets[i].BetId = common.Random10Digit()
		bets[i].Status = "PENDING"
		bets[i].Settled = false
	}

	// ---------- insert bets ----------
	if _, err = tx.NewInsert().
		Model(&bets).
		Exec(ctx); err != nil {
		return nil, err
	}

	// ---------- commit ----------
	if err = tx.Commit(); err != nil {
		return nil, err
	}

	// ---------- attach bets ----------
	transaction.Bets = bets
	return transaction, nil
}

func (s *TransactionService) holdWalletTx(
	ctx context.Context,
	tx bun.Tx,
	playerNo int64,
	amount int64,
) error {

	res, err := tx.NewUpdate().
		Model((*models.Player)(nil)).
		Set("wallet = wallet - ?", amount).
		Set("wallet_locked = wallet_locked + ?", amount).
		Where("player_no = ?", playerNo).
		Where("locked = FALSE").
		Where("wallet >= ?", amount).
		Exec(ctx)

	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return errors.New("insufficient balance or player locked")
	}

	return nil
}

func (s *TransactionService) WinBet(ctx context.Context, playerID uuid.UUID, stake int64, winAmount int64) error {

	return s.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {

		res, err := tx.NewUpdate().
			Model((*models.Player)(nil)).
			Set("wallet_locked = wallet_locked - ?", stake).
			Set("wallet = wallet + ?", winAmount).
			Set("updated_at = now()").
			Where("id = ?", playerID).
			Where("wallet_locked >= ?", stake).
			Exec(ctx)

		if err != nil {
			return err
		}

		rows, _ := res.RowsAffected()
		if rows == 0 {
			return errors.New("invalid wallet_locked or player not found")
		}

		return nil
	})
}

func (s *TransactionService) LoseBet(ctx context.Context, playerId uuid.UUID, stake int64) error {

	return s.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {

		res, err := tx.NewUpdate().
			Model((*models.Player)(nil)).
			Set("wallet_locked = wallet_locked - ?", stake).
			Set("updated_at = now()").
			Where("id = ?", playerId).
			Where("wallet_locked >= ?", stake).
			Exec(ctx)
		if err != nil {
			return err
		}
		rows, _ := res.RowsAffected()
		if rows == 0 {
			return errors.New("invalid LoseBet processing player not found")
		}

		return nil

	})
}
