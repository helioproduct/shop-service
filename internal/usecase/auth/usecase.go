package auth

import (
	"context"
	"merch-service/internal/domain"
)

type (
	SessionRepository interface {
		AddSession(context.Context, *domain.Session) error
		CheckSession(context.Context, *domain.Session) error
	}

	UserRepository interface {
		CreateUser(context.Context, string, string) (*domain.User, error)
	}
)

type AuthUsecase struct {
}

func NewAuthUsecase() *AuthUsecase {
	return &AuthUsecase{}
}
