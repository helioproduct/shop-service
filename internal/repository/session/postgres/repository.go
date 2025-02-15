package postgres

import "shop-service/config"

type SessionRepository struct {
	cfg *config.Config
}

func NewSessionRepository(cfg *config.Config) *SessionRepository {
	return &SessionRepository{
		cfg: cfg,
	}
}
