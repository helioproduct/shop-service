package middleware

import (
	"context"
	"shop-service/internal/domain"
)

type AuthUsecase interface {
	CheckSession(ctx context.Context, session *domain.Session) error
}

type AuthMiddleware struct {
	authUsecase AuthUsecase
}
