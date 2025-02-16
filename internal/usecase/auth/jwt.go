package auth

import (
	"fmt"
	"shop-service/internal/domain"
	"shop-service/pkg/logger"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   domain.UserID `json:"user_id"`
	Username string        `json:"username"`
	jwt.RegisteredClaims
}

func (uc *AuthUsecase) generateJWT(user *domain.User) (string, error) {
	expirationTime := time.Now().Add(time.Duration(uc.cfg.JWTConfig.ExpirationHours) * time.Hour)

	claims := &Claims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(uc.cfg.JWTConfig.Secret))
}

func (uc *AuthUsecase) parseJWT(tokenString string) (*Claims, error) {
	caller := "AuthUsecase.parseJWT"

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(uc.cfg.JWTConfig.Secret), nil
	})

	if err != nil {
		err = fmt.Errorf("failed to parse token: %w", err)
		logger.Error(err, caller)
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		err = fmt.Errorf("invalid token")
		logger.Error(err, caller)
		return nil, err
	}

	return claims, nil
}
