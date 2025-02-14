package inmemory

import (
	"merch-service/config"
	"merch-service/internal/domain"
	"sync"
)

type SessionRepository struct {
	cfg      *config.Config
	sessions map[domain.Session]struct{}
	mu       sync.Mutex
}

func NewSessionRepository(cfg *config.Config) *SessionRepository {
	return &SessionRepository{
		cfg: cfg,
	}
}
