package handlers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"shop-service/internal/domain"
	"shop-service/internal/usecase/auth"

	handlers "shop-service/internal/handler/auth"
	authMocks "shop-service/internal/mocks/usecase/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gopkg.in/go-playground/assert.v1"
)

func TestRegisterHandler(t *testing.T) {
	mockAuthUsecase := authMocks.NewAuthUsecase(t)
	handler := handlers.NewAuthHandlers(mockAuthUsecase)
	app := fiber.New()
	app.Post("/register", handler.Register)

	t.Run("success", func(t *testing.T) {
		mockAuthUsecase.
			On("Register", mock.Anything, auth.RegisterRequest{
				Username: "alice",
				Password: "password123",
			}).
			Return(&domain.Session{Token: "mocked-token"}, nil).
			Once()

		requestBody := `{"username":"alice","password":"password123"}`
		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(requestBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)

		var responseBody map[string]string
		err = json.Unmarshal(body, &responseBody)
		require.NoError(t, err)

		assert.Equal(t, "mocked-token", responseBody["token"])
	})

	t.Run("invalid request body", func(t *testing.T) {
		// Отправляем некорректное тело
		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(`{"username":""}`))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		io.ReadAll(resp.Body)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("usecase error", func(t *testing.T) {
		// Мок ошибки из usecase
		mockAuthUsecase.
			On("Register", mock.Anything, auth.RegisterRequest{
				Username: "alice",
				Password: "password123",
			}).
			Return(nil, domain.ErrInternalError).
			Once()

		requestBody := `{"username":"alice","password":"password123"}`
		req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(requestBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		io.ReadAll(resp.Body)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})
}
