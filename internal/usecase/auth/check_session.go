package auth

import (
	"context"
	"shop-service/internal/domain"
	"shop-service/pkg/logger"
	"time"
)

func (uc *AuthUsecase) CheckSession(ctx context.Context, session *domain.Session) error {
	caller := "AuthUsecase.CheckSession"

	claims, err := uc.parseJWT(session.Token)
	if err != nil {
		logger.Error(err, caller)
		return domain.ErrInvalidToken
	}

	session.UserID = claims.UserID
	session.Username = claims.Username

	if time.Now().After(claims.ExpiresAt.Time) {
		return domain.ErrExpiredToken
	}

	return nil
}
