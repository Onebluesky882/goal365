package bets

import (
	"context"
	analytic_module "mytipster/models/analytic"
	mybets_models "mytipster/models/mybets"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

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
		{Handicap: "-0.5", Team: "Arsenal", Odds: "1.85", Stake: "100"},
		{Handicap: "AWAY", Team: "Chelsea", Odds: "2.10", Stake: "50"},
	}

	// --- sqlmock expectation ---
	mock.ExpectQuery(`INSERT INTO "my-bets"`).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).
				AddRow(uuid.New()).
				AddRow(uuid.New()),
		)

	// --- execute service ---
	err = bets.InsertPicked(items, analyticsID, bunDB, ctx)
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

	bets, err := bets.GetBetListsByDate("2026-01-02", analyticsItems, bunDB, ctx)
	require.NoError(t, err)
	require.Len(t, bets, 1)
	require.Equal(t, analyticsID, bets[0].TipsAnalyticsID)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateMyBets_Success(t *testing.T) {
	sqlDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	require.NoError(t, err)
	defer sqlDB.Close()

	bunDB := bun.NewDB(sqlDB, pgdialect.New())

	testID := "70c39d3a-7517-43ca-8bb7-ce5929456b2c"

	body := mybets_models.BetPickIn{
		Handicap: "0.5",
		Team:     "Portadown",
		Odds:     "1.88",
		Stake:    "1000",
		Result:   "win",
		Amount: 0,
Profit : 0,
	}

	// Since Bun is inlining values, don't expect WithArgs
	mock.ExpectExec(`UPDATE "my-bets" SET bet_pick = bet_pick \|\| '.*'::jsonb WHERE \(id = '.*'\)`).
		WithoutArgs().
		WillReturnResult(sqlmock.NewResult(0, 1))

	ctx := context.Background()
	err = mybets.UpdateMyBets(testID, body, bunDB, ctx)

	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}
