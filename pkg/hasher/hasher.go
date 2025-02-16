package hasher

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

var (
	ErrHashingPassword = errors.New("failed to hash password")
)

func HashPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}

func HashAndCompare(password, hashedPassword string) bool {
	return HashPassword(password) == hashedPassword
}
