package auth

import (
	"context"
	"merch-service/internal/domain"
)

type (
	SessionRepository interface {
		PutSession(context.Context, *domain.Session) error
		CheckSession(context.Context, *domain.Session) error
	}

	UserRepository interface {
		CreateUser(context.Context, string, string) (*domain.User, error)
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
