package postgres

import (
	"context"
	"shop-service/internal/domain"
)

func (repo *SessionRepository) CheckSession(ctx context.Context, session *domain.Session) error {
	return nil
}
