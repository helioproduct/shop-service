package auth

import (
	"context"
	"fmt"
	"shop-service/internal/domain"
	userRepository "shop-service/internal/repository/user"
	"shop-service/pkg/hasher"
	"shop-service/pkg/logger"
	"time"
)

type RegisterRequest struct {
	Username string
	Password string
}

func (uc *AuthUsecase) Register(ctx context.Context, req RegisterRequest) (*domain.Session, error) {
	caller := "AuthUsecase.Register"

	createRequest, err := req.mapRegisterRequest()
	if err != nil {
		logger.Error(err, caller)
		return nil, err
	}

	user, err := uc.userRepo.CreateUser(ctx, *createRequest)
	if err != nil {
		err = fmt.Errorf("failed to create user: %w", err)
		logger.Error(err, caller)
		return nil, err
	}

	token, err := uc.generateJWT(user)
	if err != nil {
		err = fmt.Errorf("failed to generate token: %w", err)
		logger.Error(err, caller)
		return nil, err
	}

	return &domain.Session{
		UserID:    user.ID,
		Username:  user.Username,
		Token:     token,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(time.Duration(uc.cfg.JWTConfig.ExpirationHours) * time.Hour),
	}, nil
}

func (r *RegisterRequest) mapRegisterRequest() (*userRepository.CreateUserRequest, error) {
	caller := "RegisterRequest.mapRegisterRequest"

	hashedPassword, err := hasher.HashPassword(r.Password)
	if err != nil {
		err = fmt.Errorf("failed to hash password: %w", err)
		logger.Error(err, caller)
		return nil, err
	}

	return &userRepository.CreateUserRequest{
		Username:       r.Username,
		HashedPassword: hashedPassword,
	}, nil
}
