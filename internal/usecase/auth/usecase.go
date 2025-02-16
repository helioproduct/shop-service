package auth

import (
	"context"
	"shop-service/config"
	"shop-service/internal/domain"
	userRepository "shop-service/internal/repository/user"
)

type (
	UserRepository interface {
		CreateUser(ctx context.Context, req userRepository.CreateUserRequest) (*domain.User, error)
		GetUserByUsername(ctx context.Context, username string) (*domain.User, error)
		GetUserHashedPassword(ctx context.Context, username string) (string, error)
	}
)

type AuthUsecase struct {
	cfg      *config.Config
	userRepo UserRepository
}

func NewAuthUsecase(cfg *config.Config, userRepo UserRepository) *AuthUsecase {
	return &AuthUsecase{
		cfg:      cfg,
		userRepo: userRepo,
	}
}
