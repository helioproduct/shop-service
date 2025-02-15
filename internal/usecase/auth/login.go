package auth

import (
	"context"
	"shop-service/internal/domain"
)

func (uc *AuthUsecase) Login(ctx context.Context, username, password string) (*domain.Session, error) {
	return nil, nil
}
