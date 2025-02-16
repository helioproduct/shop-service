package purchase

import (
	"shop-service/internal/domain"
	"shop-service/internal/middleware"
	"shop-service/pkg/logger"

	purchaseUsecase "shop-service/internal/usecase/purchase"

	"github.com/gofiber/fiber/v2"
)

type BuyItemRequest struct {
	ProductName string `json:"productName" validate:"required"`
	// Quantity    uint64 `json:"quantity" validate:"required,min=1"`
}

type BuyItemResponse struct {
	Message string `json:"message"`
	Balance uint64 `json:"balance"`
}

func (h *Handler) HandlePurchase(c *fiber.Ctx) error {
	caller := "Handler.HandleBuyItem"

	session, ok := middleware.GetSessionFromContext(c)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	item := c.Params("item")
	if item == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "item is required"})
	}

	logger.Info(caller, "Attempting to buy item", map[string]interface{}{
		"username": session.Username,
		"item":     item,
	})

	req := BuyItemRequest{
		ProductName: item,
	}

	err := h.purchaseUsecase.BuyItemByName(c.Context(), req.mapBuyRequest(session))
	if err != nil {
		logger.Error(err, caller)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	balance, err := h.userUsecase.GetBalance(c.Context(), session.Username)
	if err != nil {
		logger.Error(err, caller)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get balance"})
	}

	return c.Status(fiber.StatusOK).JSON(BuyItemResponse{
		Message: "Item purchased successfully",
		Balance: balance,
	})
}

func (req *BuyItemRequest) mapBuyRequest(session *domain.Session) purchaseUsecase.BuyItemRequest {
	return purchaseUsecase.BuyItemRequest{
		Username:    session.Username,
		ProductName: req.ProductName,
		Quantity:    1,
	}
}
