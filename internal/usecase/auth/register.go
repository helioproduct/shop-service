package auth

import (
	"context"
	"fmt"
	"shop-service/internal/domain"
	"shop-service/pkg/hasher"
	"time"
)

type RegisterRequest struct {
	Username string
	Password string
}

func (uc *AuthUsecase) Register(ctx context.Context, req RegisterRequest) (*domain.Session, error) {
	hashedPassword, err := hasher.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user, err := uc.userRepo.CreateUser(ctx, req.Username, hashedPassword)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	token, err := uc.generateJWT(user)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &domain.Session{
		UserID:    user.ID,
		Username:  user.Username,
		Token:     token,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(time.Duration(uc.cfg.JWTConfig.ExpirationHours) * time.Hour),
	}, nil
}
