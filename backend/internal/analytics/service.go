package analytics

import (
	"context"
	"fmt"
	"log"
	m "mytipster/models/analytic"

	"github.com/uptrace/bun"
)

// AnalyticService คือ interface (สัญญา / abstraction) abstract class
type AnalyticService interface {
	InsertManual(ctx context.Context, item *m.MyAnalytics) error
	InsertMany(ctx context.Context, items []m.MyAnalytics) error
	PredictionByDay(ctx context.Context, date string) ([]m.MyAnalytics, error)
}

// constructor / factory
// (เชื่อม interface ↔ struct) คือค่า interface
func NewAnalyticService(db *bun.DB) AnalyticService {
	return &analyticsService{
		db: db,
	}
}

// inform receiver type
// คือ concrete implementation
type analyticsService struct {
	db *bun.DB
}

func (s *analyticsService) InsertManual(ctx context.Context, item *m.MyAnalytics) error {
	_, err := s.db.NewInsert().Model(item).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *analyticsService) InsertMany(ctx context.Context, items []m.MyAnalytics) error {
	var filtered []m.MyAnalytics

	for _, item := range items {
		if item.FormLeagueHomeCount < 5 {
			continue
		}
		filtered = append(filtered, item)
	}

	if len(filtered) == 0 {
		fmt.Println("No records passed the filter, nothing to insert")
		return nil
	}

	_, err := s.db.NewInsert().Model(&filtered).Exec(ctx)
	if err != nil {
		log.Fatalf("insert many error %v", err)
	}

	fmt.Printf("✅ Inserted %d records\n", len(filtered))
	return nil
}

func (s *analyticsService) PredictionByDay(ctx context.Context, date string) ([]m.MyAnalytics, error) {
	var result []m.MyAnalytics
	err := s.db.NewSelect().Model(&result).Where("date = ?", date).Scan(ctx)
	if err != nil {
		log.Fatalf("query error: %v", err)
	}
	return result, nil
}
