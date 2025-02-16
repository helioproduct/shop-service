package auth_test

import (
	"errors"
	"shop-service/config"
	"shop-service/internal/domain"
	"shop-service/internal/usecase/auth"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Claims struct {
	UserID   uint64 `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func TestGenerateJWT(t *testing.T) {

	cfg := &config.Config{
		JWTConfig: config.JWTConfig{
			Secret:          "test-secret",
			ExpirationHours: 24,
		},
	}

	user := &domain.User{
		ID:       1,
		Username: "alice",
	}

	t.Run("success", func(t *testing.T) {
		tokenString, err := auth.GenerateJWT(cfg, user)
		require.NoError(t, err)
		assert.NotEmpty(t, tokenString)

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWTConfig.Secret), nil
		})
		require.NoError(t, err)
		require.True(t, token.Valid)

		assert.Equal(t, user.ID, domain.UserID(claims.UserID))
		assert.Equal(t, user.Username, claims.Username)
		assert.WithinDuration(t, time.Now().Add(24*time.Hour), claims.ExpiresAt.Time, time.Second)
	})

	t.Run("signing error", func(t *testing.T) {
		originalGenerateJWT := auth.GenerateJWT
		defer func() { auth.GenerateJWT = originalGenerateJWT }()
		auth.GenerateJWT = func(cfg *config.Config, user *domain.User) (string, error) {
			return "", errors.New("signing failed")
		}

		tokenString, err := auth.GenerateJWT(cfg, user)
		require.Error(t, err)
		assert.Empty(t, tokenString)
		assert.Contains(t, err.Error(), "signing failed")
	})
}
