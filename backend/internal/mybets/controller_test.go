package mybets_test

import (
	"bytes"
	"encoding/json"
	"mytipster/internal/mybets"
	mybets_models "mytipster/models/mybets"
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
	reqBody := mybets_models.InsertPickedRequest{
		Items: []mybets_models.BetPickIn{
			{Picked: "HOME", Team: "Arsenal", Odds: "1.85", Stake: "100"},
			{Picked: "AWAY", Team: "Chelsea", Odds: "2.10", Stake: "50"},
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
	app.Post("/insert", mybets.InsertPickedHandler(bunDB))

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
