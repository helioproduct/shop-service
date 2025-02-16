package transfer_test

import (
	"context"
	"errors"
	"shop-service/internal/domain"
	"shop-service/internal/repository/transfer"

	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	trmsql "github.com/avito-tech/go-transaction-manager/drivers/sql/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransferRepository_CreateTransfer(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := transfer.NewTransferRepository(db, trmsql.DefaultCtxGetter)

	t.Run("Success", func(t *testing.T) {
		transfer := domain.Transfer{
			From:   1,
			To:     2,
			Amount: 100,
		}

		createdAt := time.Now()

		mock.ExpectQuery(`INSERT INTO transfers`).
			WithArgs(transfer.From, transfer.To, transfer.Amount).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at"}).
				AddRow("123e4567-e89b-12d3-a456-426614174000", createdAt))

		result, err := repo.CreateTransfer(ctx, transfer)
		require.NoError(t, err)
		require.NotNil(t, result)

		assert.Equal(t, domain.TransferID("123e4567-e89b-12d3-a456-426614174000"), result.ID)
		assert.Equal(t, transfer.From, result.From)
		assert.Equal(t, transfer.To, result.To)
		assert.Equal(t, transfer.Amount, result.Amount)
		assert.Equal(t, createdAt, result.Time)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Database Error", func(t *testing.T) {
		transfer := domain.Transfer{
			From:   1,
			To:     2,
			Amount: 200,
		}

		mock.ExpectQuery(`INSERT INTO transfers`).
			WithArgs(transfer.From, transfer.To, transfer.Amount).
			WillReturnError(errors.New("db error"))

		result, err := repo.CreateTransfer(ctx, transfer)
		require.Error(t, err)
		assert.Nil(t, result)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Scan Error (недостаточно полей)", func(t *testing.T) {
		transfer := domain.Transfer{
			From:   3,
			To:     4,
			Amount: 300,
		}

		mock.ExpectQuery(`INSERT INTO transfers`).
			WithArgs(transfer.From, transfer.To, transfer.Amount).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).
				AddRow("invalid_uuid"))

		result, err := repo.CreateTransfer(ctx, transfer)
		require.Error(t, err)
		assert.Nil(t, result)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Scan Error (неверный тип)", func(t *testing.T) {
		transfer := domain.Transfer{
			From:   5,
			To:     6,
			Amount: 500,
		}

		mock.ExpectQuery(`INSERT INTO transfers`).
			WithArgs(transfer.From, transfer.To, transfer.Amount).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at"}).
				AddRow("valid_uuid", "not_a_timestamp"))

		result, err := repo.CreateTransfer(ctx, transfer)
		require.Error(t, err)
		assert.Nil(t, result)

		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
