package transfer_test

import (
	"context"
	"errors"
	"shop-service/internal/repository/transfer"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	trmsql "github.com/avito-tech/go-transaction-manager/drivers/sql/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransferRepository_GetSentCoinsSummary(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := transfer.NewTransferRepository(db, trmsql.DefaultCtxGetter)

	fromUsername := "alice"

	t.Run("Success", func(t *testing.T) {
		mock.ExpectQuery(`SELECT tu.username AS to_username, SUM\(t.amount\) AS total_sent`).
			WithArgs(fromUsername).
			WillReturnRows(sqlmock.NewRows([]string{"to_username", "total_sent"}).
				AddRow("bob", 150).
				AddRow("carol", 200))

		results, err := repo.GetSentCoinsSummary(ctx, fromUsername)
		require.NoError(t, err)
		require.Len(t, results, 2)

		assert.Equal(t, "bob", results[0].ToUsername)
		assert.Equal(t, uint64(150), results[0].TotalSent)

		assert.Equal(t, "carol", results[1].ToUsername)
		assert.Equal(t, uint64(200), results[1].TotalSent)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Database error", func(t *testing.T) {
		mock.ExpectQuery(`SELECT tu.username AS to_username, SUM\(t.amount\) AS total_sent`).
			WithArgs(fromUsername).
			WillReturnError(errors.New("database error"))

		results, err := repo.GetSentCoinsSummary(ctx, fromUsername)
		require.Error(t, err)
		assert.Nil(t, results)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Scan error", func(t *testing.T) {
		mock.ExpectQuery(`SELECT tu.username AS to_username, SUM\(t.amount\) AS total_sent`).
			WithArgs(fromUsername).
			WillReturnRows(sqlmock.NewRows([]string{"to_username"}).AddRow("invalid"))

		results, err := repo.GetSentCoinsSummary(ctx, fromUsername)
		require.Error(t, err)
		assert.Nil(t, results)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("No results", func(t *testing.T) {
		mock.ExpectQuery(`SELECT tu.username AS to_username, SUM\(t.amount\) AS total_sent`).
			WithArgs(fromUsername).
			WillReturnRows(sqlmock.NewRows([]string{"to_username", "total_sent"}))

		results, err := repo.GetSentCoinsSummary(ctx, fromUsername)
		require.NoError(t, err)
		assert.Len(t, results, 0)

		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestTransferRepository_GetReceivedCoinsSummary(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := transfer.NewTransferRepository(db, trmsql.DefaultCtxGetter)

	toUsername := "alice"

	t.Run("Success", func(t *testing.T) {
		mock.ExpectQuery(`SELECT fu.username AS from_username, SUM\(t.amount\) AS total_received`).
			WithArgs(toUsername).
			WillReturnRows(sqlmock.NewRows([]string{"from_username", "total_received"}).
				AddRow("bob", 150).
				AddRow("carol", 300))

		results, err := repo.GetReceivedCoinsSummary(ctx, toUsername)
		require.NoError(t, err)
		require.Len(t, results, 2)

		assert.Equal(t, "bob", results[0].FromUsername)
		assert.Equal(t, uint64(150), results[0].TotalReceived)

		assert.Equal(t, "carol", results[1].FromUsername)
		assert.Equal(t, uint64(300), results[1].TotalReceived)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Database error", func(t *testing.T) {
		mock.ExpectQuery(`SELECT fu.username AS from_username, SUM\(t.amount\) AS total_received`).
			WithArgs(toUsername).
			WillReturnError(errors.New("database error"))

		results, err := repo.GetReceivedCoinsSummary(ctx, toUsername)
		require.Error(t, err)
		assert.Nil(t, results)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Scan error", func(t *testing.T) {
		mock.ExpectQuery(`SELECT fu.username AS from_username, SUM\(t.amount\) AS total_received`).
			WithArgs(toUsername).
			WillReturnRows(sqlmock.NewRows([]string{"from_username"}).AddRow("invalid"))

		results, err := repo.GetReceivedCoinsSummary(ctx, toUsername)
		require.Error(t, err)
		assert.Nil(t, results)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("No results", func(t *testing.T) {
		mock.ExpectQuery(`SELECT fu.username AS from_username, SUM\(t.amount\) AS total_received`).
			WithArgs(toUsername).
			WillReturnRows(sqlmock.NewRows([]string{"from_username", "total_received"}))

		results, err := repo.GetReceivedCoinsSummary(ctx, toUsername)
		require.NoError(t, err)
		assert.Len(t, results, 0)

		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
