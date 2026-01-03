package mybets_test

import (
	"mytipster/internal/mybets"
	m "mytipster/models/mytips"
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

	items := []m.TipsDaily{
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
	items := []m.TipsDaily{
		{ID: id},
	}

	item, err := mybets.FindId(id.String(), items)

	require.NoError(t, err)
	require.NotNil(t, item)
	assert.Equal(t, id, item.ID)
}

func TestFindId_NotFound(t *testing.T) {
	items := []m.TipsDaily{
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
	items := []m.TipsDaily{
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
	items := []m.TipsDaily{
		{ID: id},
	}

	mock.ExpectExec(`DELETE FROM "tips-daily"`).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = mybets.DeletePicked(id.String(), items, bunDB)

	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}
