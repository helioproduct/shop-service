package auth

import (
	"context"
	"shop-service/config"
	"shop-service/internal/domain"
)

type (
	UserRepository interface {
		CreateUser(ctx context.Context, username string, password string) (*domain.User, error)
		GetUserByUsername(ctx context.Context, username string) (*domain.User, error)
		GetUserHashedPassword(ctx context.Context, username string) (string, error)
	}
)

type AuthUsecase struct {
	cfg      *config.Config
	userRepo UserRepository
}

func NewAuthUsecase(userRepo UserRepository) *AuthUsecase {
	return &AuthUsecase{
		userRepo: userRepo,
	}
}
