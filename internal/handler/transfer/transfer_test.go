package transfer_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"shop-service/internal/domain"
	transferHandlers "shop-service/internal/handler/transfer"
	transferMock "shop-service/internal/mocks/usecase/transfer"
	userMock "shop-service/internal/mocks/usecase/user"
	transferUsecase "shop-service/internal/usecase/transfer"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandleTransfer(t *testing.T) {
	app := fiber.New()

	mockTransferUsecase := transferMock.NewTransferUsecsae(t)
	mockUserUsecase := userMock.NewUserUsecase(t)
	handler := transferHandlers.NewTransferHandler(mockTransferUsecase, mockUserUsecase)

	app.Post("/api/transfer", func(c *fiber.Ctx) error {
		session := &domain.Session{
			Username: "alice",
			UserID:   1,
		}
		c.Locals("session", session)
		return handler.HandleTransfer(c)
	})

	t.Run("success", func(t *testing.T) {
		reqBody := transferHandlers.TransferRequest{
			ToUsername: "bob",
			Amount:     100,
		}
		reqJSON, _ := json.Marshal(reqBody)

		mockTransferUsecase.On("SendCoins", mock.Anything, transferUsecase.SendCoinsRequest{
			From:   "alice",
			To:     "bob",
			Amount: 100,
		}).Return(nil).Once()

		mockUserUsecase.On("GetBalance", mock.Anything, "alice").
			Return(uint64(900), nil).
			Once()

		req := httptest.NewRequest(http.MethodPost, "/api/transfer", bytes.NewReader(reqJSON))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		req.Context()
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response transferHandlers.TransferResponse
		err = json.NewDecoder(resp.Body).Decode(&response)
		require.NoError(t, err)

		assert.Equal(t, "Transfer completed successfully", response.Message)
		assert.Equal(t, uint64(900), response.Balance)
	})

	// // ✅ Некорректное тело запроса
	t.Run("invalid request body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api/transfer", bytes.NewReader([]byte(`invalid-json`)))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		var response transferHandlers.ErrorResponse
		_ = json.NewDecoder(resp.Body).Decode(&response)
		assert.Equal(t, "invalid request", response.Error)
	})

	t.Run("validation error - empty username", func(t *testing.T) {
		reqBody := transferHandlers.TransferRequest{
			ToUsername: "",
			Amount:     100,
		}
		reqJSON, _ := json.Marshal(reqBody)

		req := httptest.NewRequest(http.MethodPost, "/api/transfer", bytes.NewReader(reqJSON))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		var response transferHandlers.ErrorResponse
		_ = json.NewDecoder(resp.Body).Decode(&response)
		assert.Equal(t, "invalid request", response.Error)
	})

	t.Run("validation error - zero amount", func(t *testing.T) {
		reqBody := transferHandlers.TransferRequest{
			ToUsername: "bob",
			Amount:     0,
		}
		reqJSON, _ := json.Marshal(reqBody)

		req := httptest.NewRequest(http.MethodPost, "/api/transfer", bytes.NewReader(reqJSON))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		var response transferHandlers.ErrorResponse
		_ = json.NewDecoder(resp.Body).Decode(&response)
		assert.Equal(t, "invalid request", response.Error)
	})

	t.Run("usecase error - transfer failed", func(t *testing.T) {
		reqBody := transferHandlers.TransferRequest{
			ToUsername: "bob",
			Amount:     100,
		}
		reqJSON, _ := json.Marshal(reqBody)

		mockTransferUsecase.On("SendCoins", mock.Anything, transferUsecase.SendCoinsRequest{
			From:   "alice",
			To:     "bob",
			Amount: 100,
		}).Return(errors.New("transfer failed")).Once()

		req := httptest.NewRequest(http.MethodPost, "/api/transfer", bytes.NewReader(reqJSON))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

		var response transferHandlers.ErrorResponse
		_ = json.NewDecoder(resp.Body).Decode(&response)
		assert.Equal(t, "transfer failed", response.Error)
	})

	t.Run("usecase error - failed to get balance", func(t *testing.T) {
		reqBody := transferHandlers.TransferRequest{
			ToUsername: "bob",
			Amount:     100,
		}
		reqJSON, _ := json.Marshal(reqBody)

		mockTransferUsecase.On("SendCoins", mock.Anything, transferUsecase.SendCoinsRequest{
			From:   "alice",
			To:     "bob",
			Amount: 100,
		}).Return(nil).Once()

		mockUserUsecase.On("GetBalance", mock.Anything, "alice").
			Return(uint64(0), errors.New("failed to get balance")).
			Once()

		req := httptest.NewRequest(http.MethodPost, "/api/transfer", bytes.NewReader(reqJSON))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

		var response transferHandlers.ErrorResponse
		_ = json.NewDecoder(resp.Body).Decode(&response)
		assert.Equal(t, "failed to get balance", response.Error)
	})
}
