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

	logger.Print(caller, "hashedpassowrd", hashedPassword)

	if err != nil {
		err = fmt.Errorf("failed to get user hashed password: %w", err)
		logger.Error(err, caller)
		return nil, err
	}

	if !hasher.HashAndCompare(req.Password, hashedPassword) {
		err = fmt.Errorf("invalid password")
		logger.Error(err, caller)
		return nil, err
	}

	token, err := uc.generateJWT(&domain.User{Username: req.Username})
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
