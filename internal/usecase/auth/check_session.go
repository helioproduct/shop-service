package auth

import (
	"context"
	"fmt"
	"shop-service/internal/domain"
	"shop-service/pkg/logger"
	"time"
)

func (uc *AuthUsecase) CheckSession(ctx context.Context, session *domain.Session) error {
	caller := "AuthUsecase.CheckSession"

	claims, err := uc.parseJWT(session.Token)
	if err != nil {
		err = fmt.Errorf("invalid token: %w", err)
		logger.Error(err, caller)
		return err
	}

	if claims.Username != session.Username || claims.UserID != session.UserID {
		err = fmt.Errorf("session data does not match token")
		logger.Error(err, caller)
		return err
	}

	if time.Now().After(claims.ExpiresAt.Time) {
		err = fmt.Errorf("token expired")
		logger.Error(err, caller)
		return err
	}

	return nil
}
