package user_test

import (
	"context"
	"errors"
	"shop-service/internal/domain"
	mocks "shop-service/internal/mocks/repository/user"
	"shop-service/internal/usecase/user"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserUsecase_GetBalance(t *testing.T) {
	ctx := context.Background()

	mockUserRepo := mocks.NewUserRepository(t)
	uc := user.NewUserUsecase(mockUserRepo)

	username := "testuser"

	t.Run("Success", func(t *testing.T) {
		expectedBalance := uint64(150)

		mockUserRepo.On("GetUserByUsername", ctx, username).
			Return(&domain.User{
				Username: username,
				Balance:  expectedBalance,
			}, nil).
			Once()

		balance, err := uc.GetBalance(ctx, username)
		require.NoError(t, err)
		assert.Equal(t, expectedBalance, balance)

		mockUserRepo.AssertExpectations(t)
	})

	t.Run("User Not Found", func(t *testing.T) {
		mockUserRepo.On("GetUserByUsername", ctx, username).
			Return(nil, domain.ErrUserNotFound).
			Once()

		balance, err := uc.GetBalance(ctx, username)
		require.ErrorIs(t, err, domain.ErrUserNotFound)
		assert.Equal(t, uint64(0), balance)

		mockUserRepo.AssertExpectations(t)
	})

	t.Run("Repository Error", func(t *testing.T) {
		mockUserRepo.On("GetUserByUsername", ctx, username).
			Return(nil, errors.New("database error")).
			Once()

		balance, err := uc.GetBalance(ctx, username)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "database error")
		assert.Equal(t, uint64(0), balance)

		mockUserRepo.AssertExpectations(t)
	})
}
