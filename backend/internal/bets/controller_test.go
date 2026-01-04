package bets_test

import (
	"bytes"
	"context"
	"encoding/json"
	analytic_module "mytipster/models/analytic"
	bets_models "mytipster/models/bets"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func TestInsertPickedHandler_Success(t *testing.T) {
	// --- setup sqlmock ---
	sqlDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer sqlDB.Close()

	bunDB := bun.NewDB(sqlDB, pgdialect.New())

	// --- prepare request payload ---
	reqBody := bets_models.InsertPickedRequest{
		Items: []bets_models.Bets{
			{Handicap: "HOME", Team: "Arsenal", Odds: "1.85", Stake: "100"},
			{Handicap: "AWAY", Team: "Chelsea", Odds: "2.10", Stake: "50"},
		},
	}

	bodyBytes, _ := json.Marshal(reqBody)

	// --- setup sqlmock expectation ---
	// Bun INSERT RETURNING uses ExpectQuery
	mock.ExpectQuery(`INSERT INTO "my-bets"`).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).
				AddRow(uuid.New()).
				AddRow(uuid.New()),
		)

	// --- create Fiber app ---
	app := fiber.New()
	app.Post("/insert", bets.InsertPickedHandler(bunDB))

	// --- create HTTP request ---
	req := httptest.NewRequest(http.MethodPost, "/insert", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	// --- execute request ---
	resp, err := app.Test(req)
	require.NoError(t, err)

	// --- assert response status code ---
	require.Equal(t, http.StatusOK, resp.StatusCode)

	// --- parse and assert response body ---
	var res map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&res)
	require.NoError(t, err)

	require.Equal(t, true, res["success"])
	require.Equal(t, float64(len(reqBody.Items)), res["inserted"])

	// --- verify all sqlmock expectations ---
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestBetsListsByDate_Success(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer sqlDB.Close()

	bunDB := bun.NewDB(sqlDB, pgdialect.New())
	ctx := context.Background()

	// --- sample analytics ---
	analyticsID := uuid.New()
	otherID := uuid.New()
	analyticsItems := []analytic_module.MyAnalytics{
		{ID: analyticsID, FixtureID: 1, Date: "2026-01-03"},
		{ID: otherID, FixtureID: 2, Date: "2026-01-02"},
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

	// --- call function ---
	bets, err := mybets.GetBetListsByDate("2026-01-03", analyticsItems, bunDB, ctx)
	require.NoError(t, err)

	// --- assert ---
	require.Len(t, bets, 1)
	require.Equal(t, analyticsID, bets[0].TipsAnalyticsID)

	// --- verify expectations ---
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateMyBetsHandler_Success(t *testing.T) {
	// --- setup sqlmock ---
	sqlDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer sqlDB.Close()

	bunDB := bun.NewDB(sqlDB, pgdialect.New())

	// --- fiber app ---
	app := fiber.New()
	app.Put("/my-bets", mybets.UpdateMyBetsHandler(bunDB))

	// --- test data ---
	betID := uuid.New()

	body := mybets_models.BetPickIn{
		Handicap: "0.5",
		Team:     "Arsenal",
		Odds:     "1.95",
		Stake:    "1000",
		Result:   "win",
	}

	bodyBytes, _ := json.Marshal(body)

	// --- mock SELECT ---
	mock.ExpectQuery(`SELECT .* FROM "my-bets"`).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id",
				"tips_analytics_id",
				"handicap",
				"bet_pick",
			}).AddRow(
				betID,
				uuid.New(),
				`{}`,
				`{}`,
			),
		)

	// --- mock UPDATE ---
	mock.ExpectExec(`UPDATE "my-bets"`).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// --- http request ---
	req := httptest.NewRequest(
		http.MethodPut,
		"/my-bets?id="+betID.String(),
		bytes.NewReader(bodyBytes),
	)
	req.Header.Set("Content-Type", "application/json")

	// --- perform request ---
	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	// --- assert response ---
	var res map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&res)
	require.NoError(t, err)

	require.Equal(t, true, res["success"])

	// --- verify db expectations ---
	require.NoError(t, mock.ExpectationsWereMet())
}
