package info

import (
	"shop-service/internal/middleware"
	"shop-service/pkg/constant"
	"shop-service/pkg/logger"

	"shop-service/internal/repository/purchase"

	"github.com/gofiber/fiber/v2"
)

type InventoryItem struct {
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}

type CoinHistoryItem struct {
	Username string `json:"username"`
	Amount   uint64 `json:"amount"`
}

type CoinHistory struct {
	Received []CoinHistoryItem `json:"received"`
	Sent     []CoinHistoryItem `json:"sent"`
}

type InfoResponse struct {
	Coins       uint64          `json:"coins"`
	Inventory   []InventoryItem `json:"inventory"`
	CoinHistory CoinHistory     `json:"coinHistory"`
}

func (h *Hanlder) HandleInfo(c *fiber.Ctx) error {
	caller := "Handler.HandleInfo"

	session, ok := middleware.GetSessionFromContext(c)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	coins, err := h.userUsecase.GetBalance(c.Context(), session.Username)
	if err != nil {
		logger.Error(err, caller)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get balance"})
	}

	inventoryItems, err := h.purchaseUsecase.GetSummary(c.Context(), purchase.PurchaseSummaryRequest{
		UserID: session.UserID,
		Limit:  constant.DefaultLimit,
		Offset: constant.DefaultOffet,
	})
	if err != nil {
		logger.Error(err, caller)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get inventory"})
	}

	var inventory []InventoryItem
	for _, item := range inventoryItems {
		inventory = append(inventory, InventoryItem{
			Type:     item.Product.Name,
			Quantity: int(item.Amount),
		})
	}

	sentCoins, err := h.transferUsecase.GetSentCoinsSummary(c.Context(), session.Username)
	if err != nil {
		logger.Error(err, caller)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get sent coins"})
	}

	receivedCoins, err := h.transferUsecase.GetReceivedCoinsSummary(c.Context(), session.Username)
	if err != nil {
		logger.Error(err, caller)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get received coins"})
	}

	coinHistory := CoinHistory{
		Received: make([]CoinHistoryItem, 0),
		Sent:     make([]CoinHistoryItem, 0),
	}

	for _, received := range receivedCoins {
		coinHistory.Received = append(coinHistory.Received, CoinHistoryItem{
			Username: received.FromUsername,
			Amount:   received.TotalReceived,
		})
	}

	for _, sent := range sentCoins {
		coinHistory.Sent = append(coinHistory.Sent, CoinHistoryItem{
			Username: sent.ToUsername,
			Amount:   sent.TotalSent,
		})
	}

	response := InfoResponse{
		Coins:       coins,
		Inventory:   inventory,
		CoinHistory: coinHistory,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
