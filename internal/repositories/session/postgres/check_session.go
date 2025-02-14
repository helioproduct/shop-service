package postgres

import (
	"context"
	"merch-service/internal/domain"
)

func (repo *SessionRepository) CheckSession(ctx context.Context, session *domain.Session) error {
	return nil
}
