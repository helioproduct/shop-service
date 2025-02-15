package postgres_test

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"shop-service/internal/domain"
	"shop-service/internal/repository/user/postgres"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserRepository_GetUserByUsername(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := postgres.NewUserRepository(db)

	// Тестовые данные
	username := "test_user"
	expectedUser := &domain.User{
		ID:       1,
		Username: username,
		Balance:  100,
	}

	args := []driver.Value{username}

	t.Run("Success", func(t *testing.T) {
		mock.ExpectQuery(`SELECT id, username, balance FROM users`).
			WithArgs(args...).
			WillReturnRows(sqlmock.NewRows([]string{"id", "username", "balance"}).
				AddRow(expectedUser.ID, expectedUser.Username, expectedUser.Balance))

		user, err := repo.GetUserByUsername(ctx, username)
		require.NoError(t, err)
		require.NotNil(t, user)

		assert.Equal(t, expectedUser.ID, user.ID)
		assert.Equal(t, expectedUser.Username, user.Username)
		assert.Equal(t, expectedUser.Balance, user.Balance)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("User Not Found", func(t *testing.T) {
		mock.ExpectQuery(`SELECT id, username, balance FROM users`).
			WithArgs(args...).
			WillReturnError(sql.ErrNoRows)

		user, err := repo.GetUserByUsername(ctx, username)
		require.Error(t, err)
		require.Nil(t, user)

		assert.ErrorIs(t, err, domain.ErrUserNotFound)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Query Error", func(t *testing.T) {
		mock.ExpectQuery(`SELECT id, username, balance FROM users`).
			WithArgs(args...).
			WillReturnError(errors.New("database error"))

		user, err := repo.GetUserByUsername(ctx, username)
		require.Error(t, err)
		require.Nil(t, user)

		assert.Contains(t, err.Error(), "failed to execute GetUserByUsername query")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
