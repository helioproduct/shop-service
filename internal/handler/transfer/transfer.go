package transfer

import (
	"shop-service/internal/domain"
	"shop-service/internal/middleware"
	"shop-service/pkg/logger"

	transferUsecase "shop-service/internal/usecase/transfer"

	"github.com/gofiber/fiber/v2"
)

type TransferRequest struct {
	ToUsername string `json:"toUsername" validate:"required"`
	Amount     uint64 `json:"amount" validate:"required,gt=0"`
}

type TransferResponse struct {
	Message string `json:"message"`
	Balance uint64 `json:"balance"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (h *Handler) HandleTransfer(c *fiber.Ctx) error {
	caller := "HandleTransfer"

	session, ok := middleware.GetSessionFromContext(c)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(ErrorResponse{"unauthorized"})
	}

	var req TransferRequest
	if err := c.BodyParser(&req); err != nil {
		logger.Error(err, caller)
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{"invalid request"})
	}

	if err := req.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{"invalid request"})
	}

	err := h.transferUsecase.SendCoins(c.Context(), req.mapTransferRequest(session))
	if err != nil {
		logger.Error(err, caller)
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Error: err.Error()})
	}

	balance, err := h.userUsecase.GetBalance(c.Context(), session.Username)
	if err != nil {
		logger.Error(err, caller)
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{"failed to get balance"})
	}

	return c.Status(fiber.StatusOK).JSON(TransferResponse{
		Message: "Transfer completed successfully",
		Balance: balance,
	})
}

func (r *TransferRequest) Validate() error {
	if r.ToUsername == "" {
		return domain.ErrUsernameRequired
	}

	if r.Amount == 0 {
		return domain.ErrZeroAmount
	}

	return nil
}

func (req *TransferRequest) mapTransferRequest(session *domain.Session) transferUsecase.SendCoinsRequest {
	return transferUsecase.SendCoinsRequest{
		From:   session.Username,
		To:     req.ToUsername,
		Amount: req.Amount,
	}
}
