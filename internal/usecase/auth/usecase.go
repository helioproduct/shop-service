package auth

import (
	"context"
	"merch-service/internal/domain"
)

type (
	SessionRepository interface {
		PutSession(ctx context.Context, session *domain.Session) error
		CheckSession(ctx context.Context, session *domain.Session) error
	}

	UserRepository interface {
		CreateUser(ctx context.Context, username string, password string) (*domain.User, error)
	}
)

type AuthUsecase struct {
	userRepo    UserRepository
	sessionRepo SessionRepository
}

func NewAuthUsecase(userRepo UserRepository, sessionRepo SessionRepository) *AuthUsecase {
	return &AuthUsecase{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
	}
}
