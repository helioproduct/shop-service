package auth

import (
	"context"
	"fmt"
	"shop-service/internal/domain"
	"shop-service/pkg/hasher"
	"shop-service/pkg/logger"
	"time"
)

type LoginRequest struct {
	Username string
	Password string
}

func (uc *AuthUsecase) Login(ctx context.Context, req LoginRequest) (*domain.Session, error) {
	caller := "AuthUsecase.Login"

	hashedPassword, err := uc.userRepo.GetUserHashedPassword(ctx, req.Username)
	if err != nil {
		logger.Error(err, caller)
		return nil, fmt.Errorf("failed to get user's hashed passowrd: %w", err)
	}

	if !hasher.CompareHashedPassword(hashedPassword, req.Password) {
		return nil, domain.ErrInvalidCredentials
	}

	token, err := generateJWT(uc.cfg, &domain.User{Username: req.Username})
	if err != nil {
		err = fmt.Errorf("failed to generate token: %w", err)
		logger.Error(err, caller)
		return nil, err
	}

	return &domain.Session{
		Username:  req.Username,
		Token:     token,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}, nil
}
