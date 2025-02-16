package purchase_test

import (
	"context"
	"shop-service/internal/domain"
	"shop-service/internal/repository/purchase"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	trmsql "github.com/avito-tech/go-transaction-manager/drivers/sql/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPurchaseRepository_CreatePurchase(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := purchase.NewPurchaseRepository(db, trmsql.DefaultCtxGetter)

	timestamp := time.Now()

	t.Run("Success", func(t *testing.T) {
		req := purchase.CreatePurchaseRequest{
			UserID:    domain.UserID(1),
			ProductID: domain.ProductID(100),
			Quantity:  2,
		}

		mock.ExpectQuery(`INSERT INTO purchases`).
			WithArgs(req.UserID, req.ProductID, req.Quantity).
			WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "product_id", "quantity", "created_at"}).
				AddRow("purchase-123", req.UserID, req.ProductID, 2, timestamp))

		result, err := repo.CreatePurchase(ctx, req)
		require.NoError(t, err)
		require.NotNil(t, result)

		assert.Equal(t, domain.PurchaseID("purchase-123"), result.ID)
		assert.Equal(t, req.UserID, result.UserID)
		assert.Equal(t, req.ProductID, result.ProductID)
		assert.Equal(t, uint64(2), result.Amount)
		assert.Equal(t, timestamp, result.Time)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Scan error", func(t *testing.T) {
		req := purchase.CreatePurchaseRequest{
			UserID:    domain.UserID(3),
			ProductID: domain.ProductID(300),
			Quantity:  1,
		}

		mock.ExpectQuery(`INSERT INTO purchases`).
			WithArgs(req.UserID, req.ProductID, req.Quantity).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("purchase-456"))

		result, err := repo.CreatePurchase(ctx, req)
		require.Error(t, err)
		assert.Nil(t, result)

		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
