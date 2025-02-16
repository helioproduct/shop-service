package hasher

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrHashingPassowrd = errors.New("failed to hash password")
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", ErrHashingPassowrd
	}
	return string(hashedPassword), nil
}

func HashAndCompare(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
