package postgres

import "merch-service/config"

type SessionRepository struct {
	cfg *config.Config
}

func NewSessionRepository(cfg *config.Config) *SessionRepository {
	return &SessionRepository{
		cfg: cfg,
	}
}
