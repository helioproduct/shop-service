package auth

import (
	"context"
	"merch-service/internal/domain"
)

func (uc *AuthUsecase) Login(ctx context.Context, username, password string) (*domain.Session, error) {
	return nil, nil
}
