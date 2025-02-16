package hasher_test

import (
	"shop-service/pkg/hasher"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	password := "mypassword"
	hash, err := hasher.HashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)

	assert.True(t, hasher.CompareHashedPassword(hash, password), "Пароль должен совпадать")
	assert.False(t, hasher.CompareHashedPassword(hash, "wrongpassword"), "Неверный пароль должен не совпадать")
}
