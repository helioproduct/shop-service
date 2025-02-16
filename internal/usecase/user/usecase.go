package user

import (
	"context"
	"shop-service/internal/domain"
)

type UserRepository interface {
	GetUserByUsername(ctx context.Context, username string) (*domain.User, error)
}

type UserUsecase struct {
	userRepo UserRepository
}

func NewUserUsecase(userRepo UserRepository) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}

}
