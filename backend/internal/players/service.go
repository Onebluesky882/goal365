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

func (s *PlayerService) LogPlayerLogin(
	ctx context.Context,
	req *models.PlayerLoginLogRequest,
	ip string,
	userAgent string,
) error {

	_, err := s.db.NewInsert().
		Model(&models.PlayerLoginLog{
			PlayerID:  req.PlayerID,
			IPAddress: ip,
			UserAgent: userAgent,
		}).
		Exec(ctx)

	return err
}

func (s *PlayerService) getPlayers(ctx context.Context, userId string) ([]models.Player, error) {
	var players []models.Player
	err := s.db.NewSelect().
		Model(&players).
		Where("user_id = ?", userId).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return players, err
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
		Exec(ctx)

	if err != nil {
		return nil, err
	}
	return &player, err
}

func (s *PlayerService) updateProfile(ctx context.Context, p *models.Player) error {
	_, err := s.db.NewUpdate().
		Model(p).
		Column("name", "bio", "updated_at").
		Where("id = ?", p.ID).
		Exec(ctx)
	return err
}

func (s *PlayerService) updateLevel(ctx context.Context, p *models.Player) error {
	_, err := s.db.NewUpdate().
		Model(p).
		Column("level", "exp", "updated_at").
		Where("id = ?", p.ID).
		Exec(ctx)
	return err
}
