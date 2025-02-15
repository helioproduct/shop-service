package postgres

import (
	"context"
	"shop-service/internal/domain"
)

func (repo *SessionRepository) PutSession(ctx context.Context, session *domain.Session) error {
	return nil
}
