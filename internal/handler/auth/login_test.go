package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"shop-service/internal/domain"
	"shop-service/internal/usecase/auth"

	handlers "shop-service/internal/handler/auth"
	authMocks "shop-service/internal/mocks/usecase/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestLoginHandler(t *testing.T) {
	mockAuthUsecase := authMocks.NewAuthUsecase(t)
	handler := handlers.NewAuthHandlers(mockAuthUsecase)
	app := fiber.New()
	app.Post("/login", handler.Login)

	t.Run("success", func(t *testing.T) {
		mockAuthUsecase.
			On("Login", mock.Anything, auth.LoginRequest{
				Username: "alice",
				Password: "password123",
			}).
			Return(&domain.Session{Token: "mocked-token"}, nil).
			Once()

		requestBody := `{"username":"alice","password":"password123"}`
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(requestBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var responseBody map[string]string
		err = json.Unmarshal(body, &responseBody)
		require.NoError(t, err)

		assert.Equal(t, "mocked-token", responseBody["token"])
	})

	t.Run("invalid request body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(`{"username":""}`))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		//nolint:errcheck
		io.ReadAll(resp.Body)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("unauthorized", func(t *testing.T) {
		mockAuthUsecase.
			On("Login", mock.Anything, auth.LoginRequest{
				Username: "alice",
				Password: "wrongpassword",
			}).
			Return(nil, errors.New("invalid credentials")).
			Once()

		requestBody := `{"username":"alice","password":"wrongpassword"}`
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(requestBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		require.NoError(t, err)
		defer resp.Body.Close()
		//nolint:errcheck
		io.ReadAll(resp.Body)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("internal server error", func(t *testing.T) {
		mockAuthUsecase.
			On("Login", mock.Anything, auth.LoginRequest{
				Username: "alice",
				Password: "password123",
			}).
			Return(nil, errors.New("unexpected error")).
			Once()

		requestBody := `{"username":"alice","password":"password123"}`
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(requestBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		//nolint:errcheck
		io.ReadAll(resp.Body)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})
}
