package postgres

import "shop-service/config"

type UserRepository struct {
	cfg *config.Config
}

func NewUserRepository(cfg *config.Config) *UserRepository {
	return &UserRepository{
		cfg: cfg,
	}
}
