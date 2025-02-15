package user_test

import (
	"context"
	"database/sql/driver"
	"errors"
	"shop-service/internal/domain"

	userRepo "shop-service/internal/repository/user"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	trmsql "github.com/avito-tech/go-transaction-manager/drivers/sql/v2"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func init() {
	zerolog.SetGlobalLevel(zerolog.Disabled)
}

func TestUserRepository_CreateUser(t *testing.T) {

	ctx := context.Background()

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := userRepo.NewUserRepository(db, trmsql.DefaultCtxGetter)

	req := &userRepo.CreateUserRequest{
		Username:       "test_user",
		HashedPassword: "hashed_password",
		Balance:        100,
	}

	args := []driver.Value{req.Username, req.HashedPassword, req.Balance}

	t.Run("Success", func(t *testing.T) {
		mock.ExpectQuery(`INSERT INTO users`).
			WithArgs(args...).
			WillReturnRows(sqlmock.NewRows([]string{"id", "username", "balance"}).
				AddRow(1, req.Username, req.Balance))

		user, err := repo.CreateUser(ctx, req)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, domain.UserID(1), user.ID)
		assert.Equal(t, req.Username, user.Username)
		assert.Equal(t, uint64(req.Balance), user.Balance)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Exec Error", func(t *testing.T) {
		mock.ExpectQuery(`INSERT INTO users`).
			WithArgs(args...).
			WillReturnError(errors.New("database error"))

		user, err := repo.CreateUser(ctx, req)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "failed to insert user")

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Scan Error", func(t *testing.T) {
		mock.ExpectQuery(`INSERT INTO users`).
			WithArgs(args...).
			WillReturnRows(sqlmock.NewRows([]string{"id", "username", "balance"}).
				AddRow("invalid_id", req.Username, req.Balance))

		user, err := repo.CreateUser(ctx, req)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "failed to insert user")

		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
