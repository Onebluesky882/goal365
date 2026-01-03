package mybets_test

import (
	"context"
	"mytipster/internal/mybets"
	analytic_module "mytipster/models/analytic"
	m "mytipster/models/analytic"
	mybets_models "mytipster/models/mybets"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func TestFilterTipsDailyByDate(t *testing.T) {
	date := "2026-01-02"

	items := []m.MyAnalytics{
		{
			FixtureID: 1,
			Date:      "1767369600", // 2026-01-02 UTC
		},
		{
			FixtureID: 2,
			Date:      "1767283200", // 2026-01-01 UTC
		},
	}

	result, err := mybets.FilterPredictionByDate(date, items)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result) != 1 {
		t.Fatalf("expected 1 item, got %d", len(result))
	}

	if result[0].FixtureID != 1 {
		t.Errorf("expected FixtureID 1, got %d", result[0].FixtureID)
	}
}

func TestFindId_Success(t *testing.T) {
	id := uuid.New()
	items := []m.MyAnalytics{
		{ID: id},
	}

	item, err := mybets.FindId(id.String(), items)

	require.NoError(t, err)
	require.NotNil(t, item)
	assert.Equal(t, id, item.ID)
}

func TestFindId_NotFound(t *testing.T) {
	items := []m.MyAnalytics{
		{ID: uuid.New()},
	}

	item, err := mybets.FindId(uuid.New().String(), items)

	require.Error(t, err)
	require.Nil(t, item)
}

func TestUpdatePicked(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	bunDB := bun.NewDB(db, pgdialect.New())

	id := uuid.New()
	items := []m.MyAnalytics{
		{ID: id},
	}

	mock.ExpectExec("UPDATE").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = mybets.UpdatePicked(id.String(), items, bunDB)

	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDeletePicked(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	bunDB := bun.NewDB(db, pgdialect.New())

	id := uuid.New()
	items := []m.MyAnalytics{
		{ID: id},
	}

	mock.ExpectExec(`DELETE FROM "tips-daily"`).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = mybets.DeletePicked(id.String(), items, bunDB)

	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestInsertPicked_WithAnalyticsID_Success(t *testing.T) {
	// --- mock db ---
	sqlDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer sqlDB.Close()
	bunDB := bun.NewDB(sqlDB, pgdialect.New())
	ctx := context.Background()

	// --- analyticsID สำหรับ FK ---
	analyticsID := uuid.New()

	// --- input ---
	items := []mybets_models.BetPickIn{
		{Picked: "HOME", Team: "Arsenal", Odds: "1.85", Stake: "100"},
		{Picked: "AWAY", Team: "Chelsea", Odds: "2.10", Stake: "50"},
	}

	// --- sqlmock expectation ---
	mock.ExpectQuery(`INSERT INTO "my-bets"`).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).
				AddRow(uuid.New()).
				AddRow(uuid.New()),
		)

	// --- execute service ---
	err = mybets.InsertPicked(items, analyticsID, bunDB, ctx)
	require.NoError(t, err)

	// --- verify db expectations ---
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetBetsListsByDate_Success(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer sqlDB.Close()

	bunDB := bun.NewDB(sqlDB, pgdialect.New())
	ctx := context.Background()

	analyticsID := uuid.New()
	analyticsItems := []analytic_module.MyAnalytics{
		{ID: analyticsID, FixtureID: 1, Date: "2026-01-02"},
		{ID: uuid.New(), FixtureID: 2, Date: "2026-01-03"},
	}

	// --- mock my-analytics query ---
	mock.ExpectQuery(`SELECT .* FROM "my-analytics"`).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "fixture_id", "date"}).
				AddRow(analyticsItems[0].ID, analyticsItems[0].FixtureID, analyticsItems[0].Date).
				AddRow(analyticsItems[1].ID, analyticsItems[1].FixtureID, analyticsItems[1].Date),
		)

	// --- mock my-bets query ---
	mock.ExpectQuery(`SELECT .* FROM "my-bets"`).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "tips_analytics_id"}).
				AddRow(uuid.New(), analyticsID),
		)

	bets, err := mybets.GetBetListsByDate("2026-01-02", analyticsItems, bunDB, ctx)
	require.NoError(t, err)
	require.Len(t, bets, 1)
	require.Equal(t, analyticsID, bets[0].TipsAnalyticsID)

	require.NoError(t, mock.ExpectationsWereMet())
}
