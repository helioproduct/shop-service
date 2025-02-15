package postgres_test

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"shop-service/internal/domain"
	"shop-service/internal/repository/product/postgres"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	zerolog.SetGlobalLevel(zerolog.Disabled)
}

func TestProductRepository_GetProductByName(t *testing.T) {
	ctx := context.Background()

	// Создаем mock для базы данных
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	// Создаем репозиторий с sqlmock
	repo := postgres.NewProductRepository(db)

	productName := "Test Product"
	expectedProduct := &domain.Product{
		ID:    1,
		Name:  productName,
		Price: 500,
	}

	args := []driver.Value{productName}

	// 1️⃣ Успешный кейс
	t.Run("Success", func(t *testing.T) {
		mock.ExpectQuery(`SELECT id, name, price FROM products`).
			WithArgs(args...).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price"}).
				AddRow(expectedProduct.ID, expectedProduct.Name, expectedProduct.Price))

		product, err := repo.GetProductByName(ctx, productName)
		require.NoError(t, err)
		require.NotNil(t, product)

		assert.Equal(t, expectedProduct.ID, product.ID)
		assert.Equal(t, expectedProduct.Name, product.Name)
		assert.Equal(t, expectedProduct.Price, product.Price)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	// 2️⃣ Товар не найден
	t.Run("Product Not Found", func(t *testing.T) {
		mock.ExpectQuery(`SELECT id, name, price FROM products`).
			WithArgs(args...).
			WillReturnError(sql.ErrNoRows)

		product, err := repo.GetProductByName(ctx, productName)
		require.ErrorIs(t, err, domain.ErrProductNotFound)
		require.Nil(t, product)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	// 3️⃣ Ошибка выполнения запроса
	t.Run("Database Error", func(t *testing.T) {
		mock.ExpectQuery(`SELECT id, name, price FROM products`).
			WithArgs(args...).
			WillReturnError(errors.New("database error"))

		product, err := repo.GetProductByName(ctx, productName)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "failed to execute GetProductByName query")
		require.Nil(t, product)

		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
