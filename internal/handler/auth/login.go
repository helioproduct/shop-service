package handlers

import (
	"github.com/gofiber/fiber/v2"

	authUsecase "shop-service/internal/usecase/auth"
	"shop-service/pkg/logger"
)

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

func (h *AuthHandlers) Login(c *fiber.Ctx) error {
	caller := "AuthHandlers.Login"

	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	if err := req.Validate(); err != nil {
		logger.Error(err, caller)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	session, err := h.authUC.Login(c.Context(), req.mapLoginRequest())
	if err != nil {
		logger.Error(err, caller)

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(AuthResponse{Token: session.Token})
}

func (r *LoginRequest) Validate() error {
	if r.Username == "" {
		return fiber.NewError(fiber.StatusBadRequest, "username required")
	}
	if r.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "password required")
	}

	return nil
}

func (r *LoginRequest) mapLoginRequest() authUsecase.LoginRequest {
	return authUsecase.LoginRequest{
		Username: r.Username,
		Password: r.Password,
	}
}
