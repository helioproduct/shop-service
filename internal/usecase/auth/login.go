package auth

import (
	"context"
	"fmt"
	"shop-service/internal/domain"
	"shop-service/pkg/hasher"
	"time"
)

func (uc *AuthUsecase) Login(ctx context.Context, username, password string) (*domain.Session, error) {
	hashedPassword, err := uc.userRepo.GetUserHashedPassword(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user hashed password: %w", err)
	}

	if !hasher.HashAndCompare(password, hashedPassword) {
		return nil, fmt.Errorf("invalid password")

	}

	token, err := uc.generateJWT(&domain.User{Username: username})
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &domain.Session{
		Username:  username,
		Token:     token,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}, nil
}
