package postgres

import (
	"context"
	"merch-service/internal/domain"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (repo *UserRepository) CreateUser(ctx context.Context, username, passord string) (*domain.User, error) {
	return nil, nil
}
