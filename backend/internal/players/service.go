package player

import (
	"context"
	"mytipster/lib/common"
	"mytipster/models"

	"github.com/uptrace/bun"
)

type PlayerService struct {
	db *bun.DB
}

func NewPlayer(db *bun.DB) *PlayerService {
	return &PlayerService{
		db: db,
	}
}

func (s *PlayerService) CreatePlayer(
	ctx context.Context,
	name string,
	userID string,
) (*models.Player, error) {

	player := models.Player{
		UserId:   userID,
		Name:     name,
		PlayerNo: common.Random10Digit(),
		Wallet:   100000,
		Level:    1,
	}

	_, err := s.db.NewInsert().
		Model(&player).
		Where("user_id = ?", userID).
		Exec(ctx)

	if err != nil {
		return nil, err
	}
	return &player, err
}
