package user_test

import (
	"context"
	"errors"
	"shop-service/internal/domain"
	userRepo "shop-service/internal/repository/user"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	trmsql "github.com/avito-tech/go-transaction-manager/drivers/sql/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserRepository_UpdateUser(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := userRepo.NewUserRepository(db, trmsql.DefaultCtxGetter)

	username := "new_username"
	hashedPassword := "new_hashed_password"
	balance := uint64(200)

	t.Run("Success - Update all fields", func(t *testing.T) {
		req := &userRepo.UpdateUserRequest{
			UserID:         1,
			Username:       &username,
			HashedPassword: &hashedPassword,
			Balance:        &balance,
		}

		mock.ExpectExec(`UPDATE users`).
			WithArgs(username, hashedPassword, balance, req.UserID).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.UpdateUser(ctx, req)
		require.NoError(t, err)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Success - Update username only", func(t *testing.T) {
		req := &userRepo.UpdateUserRequest{
			UserID:   2,
			Username: &username,
		}

		mock.ExpectExec(`UPDATE users`).
			WithArgs(username, req.UserID).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.UpdateUser(ctx, req)
		require.NoError(t, err)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("User Not Found", func(t *testing.T) {
		req := &userRepo.UpdateUserRequest{
			UserID:   999,
			Username: &username,
		}

		mock.ExpectExec(`UPDATE users`).
			WithArgs(username, req.UserID).
			WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.UpdateUser(ctx, req)
		require.ErrorIs(t, err, domain.ErrUserNotFound)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Database Error", func(t *testing.T) {
		req := &userRepo.UpdateUserRequest{
			UserID:   3,
			Username: &username,
		}

		mock.ExpectExec(`UPDATE users`).
			WithArgs(username, req.UserID).
			WillReturnError(errors.New("database error"))

		err := repo.UpdateUser(ctx, req)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "failed to execute UpdateUser query")

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Error Getting Rows Affected", func(t *testing.T) {
		req := &userRepo.UpdateUserRequest{
			UserID:   4,
			Username: &username,
		}

		mock.ExpectExec(`UPDATE users`).
			WithArgs(username, req.UserID).
			WillReturnResult(sqlmock.NewErrorResult(errors.New("rows affected error")))

		err := repo.UpdateUser(ctx, req)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "failed to get affected rows")

		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
