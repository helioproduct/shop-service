package auth

import "merch-service/internal/domain"

func (u *AuthUsecase) Register(username, password string) *domain.User {
	return &domain.User{}
}
