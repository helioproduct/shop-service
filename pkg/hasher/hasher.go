package hasher

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrHashingPassword  = errors.New("failed to hash password")
	ErrPasswordMismatch = errors.New("password does not match")
)

// HashPassword — Хеширует пароль через bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", ErrHashingPassword
	}
	return string(hashedPassword), nil
}

// CompareHashedPassword — Сравнивает хеш и пароль
func CompareHashedPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
