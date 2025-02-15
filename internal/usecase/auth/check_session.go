package auth

import (
	"context"
	"shop-service/internal/domain"
)

func (uc *AuthUsecase) CheckSession(ctx context.Context, session *domain.Session) error {
	return nil
}
