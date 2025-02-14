package postgres

import (
	"context"
	"merch-service/config"
	"merch-service/internal/domain"
)

type UserRepository struct {
	cfg *config.Config
}

func NewUserRepository(cfg *config.Config) *UserRepository {
	return &UserRepository{
		cfg: cfg,
	}
}

func (repo *UserRepository) CreateUser(ctx context.Context, username, passord string) (*domain.User, error) {
	return nil, nil
}
