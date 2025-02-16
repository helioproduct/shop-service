package transfer_test

import (
	"context"
	"errors"
	"shop-service/internal/repository/transfer"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	trmsql "github.com/avito-tech/go-transaction-manager/drivers/sql/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransferRepository_GetTransfers(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := transfer.NewTransferRepository(db, trmsql.DefaultCtxGetter)

	t.Run("Success - Filter by FromUsername", func(t *testing.T) {
		fromUsername := "alice"
		filter := transfer.GetTransfersFilter{
			FromUsername: &fromUsername,
		}

		mock.ExpectQuery(`SELECT`).
			WithArgs(fromUsername).
			WillReturnRows(sqlmock.NewRows([]string{
				"id", "from_user_id", "to_user_id", "amount", "created_at", "from_username", "to_username"}).
				AddRow("t1", 1, 2, 150, time.Now(), "alice", "bob").
				AddRow("t2", 1, 3, 200, time.Now(), "alice", "charlie"))

		results, err := repo.GetTransfers(ctx, filter)
		require.NoError(t, err)
		require.Len(t, results, 2)

		assert.Equal(t, "alice", results[0].FromUsername)
		assert.Equal(t, "bob", results[0].ToUsername)
		assert.Equal(t, uint64(150), results[0].Amount)

		assert.Equal(t, "alice", results[1].FromUsername)
		assert.Equal(t, "charlie", results[1].ToUsername)
		assert.Equal(t, uint64(200), results[1].Amount)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Success - Filter by ToUsername", func(t *testing.T) {
		toUsername := "bob"
		filter := transfer.GetTransfersFilter{
			ToUsername: &toUsername,
		}

		mock.ExpectQuery(`SELECT`).
			WithArgs(toUsername).
			WillReturnRows(sqlmock.NewRows([]string{
				"id", "from_user_id", "to_user_id", "amount", "created_at", "from_username", "to_username"}).
				AddRow("t3", 2, 3, 300, time.Now(), "carol", "bob"))

		results, err := repo.GetTransfers(ctx, filter)
		require.NoError(t, err)
		require.Len(t, results, 1)

		assert.Equal(t, "carol", results[0].FromUsername)
		assert.Equal(t, "bob", results[0].ToUsername)
		assert.Equal(t, uint64(300), results[0].Amount)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Success - Filter by FromUsername and ToUsername", func(t *testing.T) {
		fromUsername := "alice"
		toUsername := "bob"

		filter := transfer.GetTransfersFilter{
			FromUsername: &fromUsername,
			ToUsername:   &toUsername,
		}

		mock.ExpectQuery(`SELECT`).
			WithArgs(fromUsername, toUsername).
			WillReturnRows(sqlmock.NewRows([]string{
				"id", "from_user_id", "to_user_id", "amount", "created_at", "from_username", "to_username"}).
				AddRow("t4", 1, 2, 500, time.Now(), "alice", "bob"))

		results, err := repo.GetTransfers(ctx, filter)
		require.NoError(t, err)
		require.Len(t, results, 1)

		assert.Equal(t, "alice", results[0].FromUsername)
		assert.Equal(t, "bob", results[0].ToUsername)
		assert.Equal(t, uint64(500), results[0].Amount)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Database error", func(t *testing.T) {
		fromUsername := "alice"
		filter := transfer.GetTransfersFilter{
			FromUsername: &fromUsername,
		}

		mock.ExpectQuery(`SELECT`).
			WithArgs(fromUsername).
			WillReturnError(errors.New("database error"))

		results, err := repo.GetTransfers(ctx, filter)
		require.Error(t, err)
		assert.Nil(t, results)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Scan error", func(t *testing.T) {
		fromUsername := "alice"
		filter := transfer.GetTransfersFilter{
			FromUsername: &fromUsername,
		}

		mock.ExpectQuery(`SELECT`).
			WithArgs(fromUsername).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("invalid"))

		results, err := repo.GetTransfers(ctx, filter)
		require.Error(t, err)
		assert.Nil(t, results)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("No results", func(t *testing.T) {
		fromUsername := "alice"
		filter := transfer.GetTransfersFilter{
			FromUsername: &fromUsername,
		}

		mock.ExpectQuery(`SELECT`).
			WithArgs(fromUsername).
			WillReturnRows(sqlmock.NewRows([]string{
				"id", "from_user_id", "to_user_id", "amount", "created_at", "from_username", "to_username"}))

		results, err := repo.GetTransfers(ctx, filter)
		require.NoError(t, err)
		assert.Len(t, results, 0)

		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
