package hasher_test

import (
	"shop-service/pkg/hasher"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordHashingAndChecking(t *testing.T) {
	password := "supersecret"
	hashedPassword := hasher.HashPassword(password)
	assert.True(t, hasher.HashAndCompare(password, hashedPassword), "Correct password should match")
	assert.False(t, hasher.HashAndCompare("wrongpassword", hashedPassword), "Wrong password should not match")
}
