package info_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"shop-service/internal/domain"
	"shop-service/internal/handler/info"
	"shop-service/internal/repository/purchase"
	"shop-service/pkg/constant"

	purchaseMockUsecase "shop-service/internal/mocks/usecase/purchase"
	transferMockUsecase "shop-service/internal/mocks/usecase/trasnfer"
	userMockUsecase "shop-service/internal/mocks/usecase/user"

	transferRepository "shop-service/internal/repository/transfer"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandleInfo(t *testing.T) {
	app := fiber.New()

	mockTransferUsecase := transferMockUsecase.NewTransferUsecsae(t)
	mockPurchaseUsecase := purchaseMockUsecase.NewPurchaseUsecase(t)
	mockUserUsecase := userMockUsecase.NewUserUsecase(t)

	handler := info.NewInfoHandler(mockPurchaseUsecase, mockTransferUsecase, mockUserUsecase)

	app.Get("/api/info", func(c *fiber.Ctx) error {
		session := &domain.Session{
			Username: "alice",
			UserID:   1,
		}
		c.Locals("session", session)
		return handler.HandleInfo(c)
	})

	t.Run("success", func(t *testing.T) {
		mockUserUsecase.On(
			"GetBalance",
			mock.Anything,
			"alice",
		).Return(uint64(500), nil).Once()

		mockPurchaseUsecase.On(
			"GetSummary",
			mock.Anything,
			purchase.PurchaseSummaryRequest{
				UserID: 1,
				Limit:  constant.DefaultLimit,
				Offset: constant.DefaultOffet,
			},
		).Return([]*purchase.PurchaseSummary{
			{
				Product: domain.Product{Name: "Sword"},
				Amount:  2,
			},
			{
				Product: domain.Product{Name: "Shield"},
				Amount:  1,
			},
		}, nil).Once()

		mockTransferUsecase.On(
			"GetSentCoinsSummary",
			mock.Anything,
			"alice",
		).Return([]*transferRepository.SentCoinsSummary{
			{ToUsername: "bob", TotalSent: 150},
			{ToUsername: "charlie", TotalSent: 200},
		}, nil).Once()

		mockTransferUsecase.On(
			"GetReceivedCoinsSummary",
			mock.Anything,
			"alice",
		).Return([]*transferRepository.ReceivedCoinsSummary{
			{FromUsername: "bob", TotalReceived: 100},
			{FromUsername: "david", TotalReceived: 250},
		}, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/api/info", bytes.NewReader(nil))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response info.InfoResponse
		err = json.NewDecoder(resp.Body).Decode(&response)
		require.NoError(t, err)

		assert.Equal(t, uint64(500), response.Coins)
		assert.Len(t, response.Inventory, 2)
		assert.Equal(t, "Sword", response.Inventory[0].Type)
		assert.Equal(t, 2, response.Inventory[0].Quantity)
		assert.Equal(t, "Shield", response.Inventory[1].Type)
		assert.Equal(t, 1, response.Inventory[1].Quantity)

		assert.Len(t, response.CoinHistory.Sent, 2)
		assert.Equal(t, "bob", response.CoinHistory.Sent[0].Username)
		assert.Equal(t, uint64(150), response.CoinHistory.Sent[0].Amount)
		assert.Equal(t, "charlie", response.CoinHistory.Sent[1].Username)
		assert.Equal(t, uint64(200), response.CoinHistory.Sent[1].Amount)

		assert.Len(t, response.CoinHistory.Received, 2)
		assert.Equal(t, "bob", response.CoinHistory.Received[0].Username)
		assert.Equal(t, uint64(100), response.CoinHistory.Received[0].Amount)
		assert.Equal(t, "david", response.CoinHistory.Received[1].Username)
		assert.Equal(t, uint64(250), response.CoinHistory.Received[1].Amount)
	})

	t.Run("unauthorized", func(t *testing.T) {
		app.Get("/api/info/unauthorized", func(c *fiber.Ctx) error {
			return handler.HandleInfo(c)
		})

		req := httptest.NewRequest(http.MethodGet, "/api/info/unauthorized", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("error getting balance", func(t *testing.T) {
		mockUserUsecase.On(
			"GetBalance",
			mock.Anything,
			"alice",
		).Return(uint64(0), errors.New("balance error")).Once()

		req := httptest.NewRequest(http.MethodGet, "/api/info", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("error getting inventory", func(t *testing.T) {
		mockUserUsecase.On(
			"GetBalance",
			mock.Anything,
			"alice",
		).Return(uint64(500), nil).Once()

		mockPurchaseUsecase.On(
			"GetSummary",
			mock.Anything,
			purchase.PurchaseSummaryRequest{
				UserID: 1,
				Limit:  constant.DefaultLimit,
				Offset: constant.DefaultOffet,
			},
		).Return(nil, errors.New("inventory error")).Once()

		req := httptest.NewRequest(http.MethodGet, "/api/info", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("error getting sent coins", func(t *testing.T) {
		mockUserUsecase.On(
			"GetBalance",
			mock.Anything,
			"alice",
		).Return(uint64(500), nil).Once()

		mockPurchaseUsecase.On(
			"GetSummary",
			mock.Anything,
			purchase.PurchaseSummaryRequest{
				UserID: 1,
				Limit:  constant.DefaultLimit,
				Offset: constant.DefaultOffet,
			},
		).Return([]*purchase.PurchaseSummary{}, nil).Once()

		mockTransferUsecase.On(
			"GetSentCoinsSummary",
			mock.Anything,
			"alice",
		).Return(nil, errors.New("sent coins error")).Once()

		req := httptest.NewRequest(http.MethodGet, "/api/info", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("error getting received coins", func(t *testing.T) {
		mockUserUsecase.On(
			"GetBalance",
			mock.Anything,
			"alice",
		).Return(uint64(500), nil).Once()

		mockPurchaseUsecase.On(
			"GetSummary",
			mock.Anything,
			purchase.PurchaseSummaryRequest{
				UserID: 1,
				Limit:  constant.DefaultLimit,
				Offset: constant.DefaultOffet,
			},
		).Return([]*purchase.PurchaseSummary{}, nil).Once()

		mockTransferUsecase.On(
			"GetSentCoinsSummary",
			mock.Anything,
			"alice",
		).Return([]*transferRepository.SentCoinsSummary{}, nil).Once()

		mockTransferUsecase.On(
			"GetReceivedCoinsSummary",
			mock.Anything,
			"alice",
		).Return(nil, errors.New("received coins error")).Once()

		req := httptest.NewRequest(http.MethodGet, "/api/info", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})
}
