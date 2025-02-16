package auth_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"shop-service/config"
	"shop-service/internal/domain"
	mocks "shop-service/internal/mocks/repository/user"
	"shop-service/internal/usecase/auth"
	"shop-service/pkg/hasher"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthUsecase_Login(t *testing.T) {
	cfg := &config.Config{
		JWTConfig: config.JWTConfig{
			Secret:          "secret-key",
			ExpirationHours: 24,
		},
	}

	mockUserRepo := mocks.NewUserRepository(t)
	uc := auth.NewAuthUsecase(cfg, mockUserRepo)

	ctx := context.Background()
	req := auth.LoginRequest{
		Username: "alice",
		Password: "password123",
	}

	hashedPassword, _ := hasher.HashPassword("password123")

	t.Run("success", func(t *testing.T) {
		mockUserRepo.
			On("GetUserHashedPassword", ctx, req.Username).
			Return(hashedPassword, nil).
			Once()

		session, err := uc.Login(ctx, req)
		require.NoError(t, err)
		assert.NotNil(t, session)
		assert.Equal(t, req.Username, session.Username)
		assert.NotEmpty(t, session.Token)
		assert.WithinDuration(t, time.Now(), session.IssuedAt, time.Second)
		assert.WithinDuration(t, time.Now().Add(24*time.Hour), session.ExpiresAt, time.Second)
	})

	t.Run("error getting user hashed password", func(t *testing.T) {
		mockUserRepo.
			On("GetUserHashedPassword", ctx, req.Username).
			Return("", errors.New("database error")).
			Once()

		session, err := uc.Login(ctx, req)
		require.Error(t, err)
		assert.Nil(t, session)
		assert.Contains(t, err.Error(), "failed to get user's hashed passowrd")
	})

	t.Run("invalid password", func(t *testing.T) {

		hashedPassword, _ := hasher.HashPassword("wrongpassword")

		mockUserRepo.
			On("GetUserHashedPassword", ctx, req.Username).
			Return(hashedPassword, nil).
			Once()

		session, err := uc.Login(ctx, req)
		require.Error(t, err)
		assert.Nil(t, session)
		assert.ErrorIs(t, err, domain.ErrInvalidCredentials)
	})

}
