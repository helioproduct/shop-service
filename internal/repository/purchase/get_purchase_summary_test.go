package purchase_test

import (
	"context"
	"errors"
	"shop-service/internal/domain"
	"shop-service/internal/repository/purchase"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	trmsql "github.com/avito-tech/go-transaction-manager/drivers/sql/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPurchaseRepository_GetPurchaseSummary(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := purchase.NewPurchaseRepository(db, trmsql.DefaultCtxGetter)

	t.Run("Success", func(t *testing.T) {
		req := purchase.PurchaseSummaryRequest{
			UserID: domain.UserID(1),
			Limit:  10,
			Offset: 0,
		}

		mock.ExpectQuery(`SELECT pr.id AS product_id, pr.name AS product_name, pr.price, SUM\(p.quantity\) AS amount`).
			WithArgs(req.UserID).
			WillReturnRows(sqlmock.NewRows([]string{"product_id", "product_name", "price", "amount"}).
				AddRow(uint64(1001), "T-Shirt", uint64(80), uint64(3)).
				AddRow(uint64(1002), "Cup", uint64(20), uint64(5)))

		results, err := repo.GetPurchaseSummary(ctx, req)
		require.NoError(t, err)
		require.Len(t, results, 2)

		assert.Equal(t, domain.ProductID(1001), results[0].Product.ID)
		assert.Equal(t, "T-Shirt", results[0].Product.Name)
		assert.Equal(t, uint64(80), results[0].Product.Price)
		assert.Equal(t, uint64(3), results[0].Amount)

		assert.Equal(t, domain.ProductID(1002), results[1].Product.ID)
		assert.Equal(t, "Cup", results[1].Product.Name)
		assert.Equal(t, uint64(20), results[1].Product.Price)
		assert.Equal(t, uint64(5), results[1].Amount)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Database Error", func(t *testing.T) {
		req := purchase.PurchaseSummaryRequest{
			UserID: domain.UserID(2),
			Limit:  5,
			Offset: 0,
		}

		mock.ExpectQuery(`SELECT pr.id AS product_id, pr.name AS product_name, pr.price, SUM\(p.quantity\) AS amount`).
			WithArgs(req.UserID).
			WillReturnError(errors.New("database error"))

		results, err := repo.GetPurchaseSummary(ctx, req)
		require.Error(t, err)
		assert.Nil(t, results)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Scan Error", func(t *testing.T) {
		req := purchase.PurchaseSummaryRequest{
			UserID: domain.UserID(3),
			Limit:  5,
			Offset: 0,
		}

		mock.ExpectQuery(`SELECT pr.id AS product_id, pr.name AS product_name, pr.price, SUM\(p.quantity\) AS amount`).
			WithArgs(req.UserID).
			WillReturnRows(sqlmock.NewRows([]string{"product_id"}).
				AddRow(uint64(1003))) // Пропущены столбцы

		results, err := repo.GetPurchaseSummary(ctx, req)
		require.Error(t, err)
		assert.Nil(t, results)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Empty Result", func(t *testing.T) {
		req := purchase.PurchaseSummaryRequest{
			UserID: domain.UserID(4),
			Limit:  5,
			Offset: 0,
		}

		mock.ExpectQuery(`SELECT pr.id AS product_id, pr.name AS product_name, pr.price, SUM\(p.quantity\) AS amount`).
			WithArgs(req.UserID).
			WillReturnRows(sqlmock.NewRows([]string{"product_id", "product_name", "price", "amount"}))

		results, err := repo.GetPurchaseSummary(ctx, req)
		require.NoError(t, err)
		assert.Len(t, results, 0)

		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
