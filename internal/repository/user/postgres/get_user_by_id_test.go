package postgres_test

import (
	"context"
	"database/sql"
	"errors"
	"shop-service/internal/domain"
	"shop-service/internal/repository/user/postgres"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	trmsql "github.com/avito-tech/go-transaction-manager/drivers/sql/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Тест GetUserByID
func TestUserRepository_GetUserByID(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := postgres.NewUserRepository(db, trmsql.DefaultCtxGetter)

	userID := domain.UserID(42)
	username := "test_user"
	balance := uint64(100)

	t.Run("Success", func(t *testing.T) {
		mock.ExpectQuery(`SELECT id, username, balance FROM users`).
			WithArgs(userID).
			WillReturnRows(sqlmock.NewRows([]string{"id", "username", "balance"}).
				AddRow(userID, username, balance))

		user, err := repo.GetUserByID(ctx, userID)
		require.NoError(t, err)
		require.NotNil(t, user)

		assert.Equal(t, userID, user.ID)
		assert.Equal(t, username, user.Username)
		assert.Equal(t, balance, user.Balance)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("User Not Found", func(t *testing.T) {
		mock.ExpectQuery(`SELECT id, username, balance FROM users`).
			WithArgs(userID).
			WillReturnError(sql.ErrNoRows)

		user, err := repo.GetUserByID(ctx, userID)
		require.ErrorIs(t, err, domain.ErrUserNotFound)
		require.Nil(t, user)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Database Error", func(t *testing.T) {
		mock.ExpectQuery(`SELECT id, username, balance FROM users`).
			WithArgs(userID).
			WillReturnError(errors.New("database error"))

		user, err := repo.GetUserByID(ctx, userID)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "failed to execute GetUserByID query")
		require.Nil(t, user)

		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
