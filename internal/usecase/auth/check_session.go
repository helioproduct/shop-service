package auth

import (
	"context"
	"shop-service/internal/domain"
	"time"
)

func (uc *AuthUsecase) CheckSession(ctx context.Context, session *domain.Session) error {
	// caller := "AuthUsecase.CheckSession"

	claims, err := uc.parseJWT(session.Token)
	if err != nil {
		return domain.ErrInvalidToken
	}

	if claims.Username != session.Username || claims.UserID != session.UserID {
		return domain.ErrTokenSessionMismatch
	}

	if time.Now().After(claims.ExpiresAt.Time) {
		return domain.ErrExpiredToken
	}

	return nil
}
