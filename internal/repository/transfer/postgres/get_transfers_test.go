package postgres_test

import (
	"context"
	"errors"
	"shop-service/internal/domain"
	"shop-service/internal/repository/transfer/postgres"
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

	repo := postgres.NewTransferRepository(db, trmsql.DefaultCtxGetter)

	t.Run("Success with filters", func(t *testing.T) {
		fromUserID := domain.UserID(1)
		toUserID := domain.UserID(2)
		limit := uint64(5)
		offset := uint64(0)

		filter := &postgres.GetTransfersFilter{
			FromUserID: &fromUserID,
			ToUserID:   &toUserID,
			Limit:      limit,
			Offset:     offset,
		}

		mock.ExpectQuery(`SELECT id, from_user_id, to_user_id, amount, created_at FROM transfers`).
			WithArgs(fromUserID, toUserID).
			WillReturnRows(sqlmock.NewRows([]string{"id", "from_user_id", "to_user_id", "amount", "created_at"}).
				AddRow("t1", fromUserID, toUserID, 100, time.Now()).
				AddRow("t2", fromUserID, toUserID, 200, time.Now()))

		results, err := repo.GetTransfers(ctx, filter)
		require.NoError(t, err)
		require.Len(t, results, 2)

		assert.Equal(t, domain.TransferID("t1"), results[0].ID)
		assert.Equal(t, fromUserID, results[0].From)
		assert.Equal(t, toUserID, results[0].To)
		assert.Equal(t, uint64(100), results[0].Amount)

		assert.Equal(t, domain.TransferID("t2"), results[1].ID)
		assert.Equal(t, uint64(200), results[1].Amount)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Database error", func(t *testing.T) {
		fromUserID := domain.UserID(1)
		filter := &postgres.GetTransfersFilter{
			FromUserID: &fromUserID,
		}

		mock.ExpectQuery(`SELECT id, from_user_id, to_user_id, amount, created_at FROM transfers`).
			WithArgs(fromUserID).
			WillReturnError(errors.New("database error"))

		results, err := repo.GetTransfers(ctx, filter)
		require.Error(t, err)
		assert.Nil(t, results)
		assert.Contains(t, err.Error(), "failed to execute GetTransfers query")

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Scan error", func(t *testing.T) {
		fromUserID := domain.UserID(1)
		filter := &postgres.GetTransfersFilter{
			FromUserID: &fromUserID,
		}

		mock.ExpectQuery(`SELECT id, from_user_id, to_user_id, amount, created_at FROM transfers`).
			WithArgs(fromUserID).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("invalid"))

		results, err := repo.GetTransfers(ctx, filter)
		require.Error(t, err)
		assert.Nil(t, results)
		assert.Contains(t, err.Error(), "failed to scan GetTransfers result")

		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
