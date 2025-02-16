package handlers

import (
	authUsecase "shop-service/internal/usecase/auth"
	"shop-service/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (h *AuthHandlers) Register(c *fiber.Ctx) error {
	caller := "AuthHandlers.Register"

	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		logger.Error(err, caller)

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := req.Validate(); err != nil {
		logger.Error(err, caller)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	} else {
		logger.Log.Println("validate ok")
	}

	session, err := h.authUC.Register(c.Context(), req.mapRegisterRequest())
	if err != nil {
		logger.Error(err, caller)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(AuthResponse{Token: session.Token})
}

func (r *RegisterRequest) Validate() error {
	if r.Username == "" {
		return fiber.NewError(fiber.StatusBadRequest, "username required")
	}
	if r.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "password required")
	}

	return nil
}

func (r *RegisterRequest) mapRegisterRequest() authUsecase.RegisterRequest {
	return authUsecase.RegisterRequest{
		Username: r.Username,
		Password: r.Password,
	}
}
