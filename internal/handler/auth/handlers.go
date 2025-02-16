package handlers

import (
	"context"
	"shop-service/internal/domain"
)

type (
	AuthUsecase interface {
		Register(ctx context.Context, username, password string) (*domain.Session, error)
		Login(ctx context.Context, username, password string) (*domain.Session, error)
		CheckSession(ctx context.Context, session *domain.Session) error
	}
)

type AuthHandlers struct {
	authUC AuthUsecase
}

func NewAuthHandlers(authUC AuthUsecase) *AuthHandlers {
	return &AuthHandlers{
		authUC: authUC,
	}
}
