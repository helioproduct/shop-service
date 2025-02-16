package auth

import (
	"context"
	"fmt"
	"shop-service/internal/domain"
	"shop-service/pkg/hasher"
	"time"
)

type LoginRequest struct {
	Username string
	Password string
}

func (uc *AuthUsecase) Login(ctx context.Context, req LoginRequest) (*domain.Session, error) {
	hashedPassword, err := uc.userRepo.GetUserHashedPassword(ctx, req.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user hashed password: %w", err)
	}

	if !hasher.HashAndCompare(req.Password, hashedPassword) {
		return nil, fmt.Errorf("invalid password")

	}

	token, err := uc.generateJWT(&domain.User{Username: req.Username})
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &domain.Session{
		Username:  req.Username,
		Token:     token,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}, nil
}
