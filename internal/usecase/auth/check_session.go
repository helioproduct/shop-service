package auth

import (
	"context"
	"fmt"
	"shop-service/internal/domain"
	"time"
)

func (uc *AuthUsecase) CheckSession(ctx context.Context, session *domain.Session) error {
	claims, err := uc.parseJWT(session.Token)
	if err != nil {
		return fmt.Errorf("invalid token: %w", err)
	}

	if claims.Username != session.Username || claims.UserID != session.UserID {
		return fmt.Errorf("session data does not match token")
	}

	if time.Now().After(claims.ExpiresAt.Time) {
		return fmt.Errorf("token expired")
	}

	return nil
}
