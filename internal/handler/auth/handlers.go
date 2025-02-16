package handlers

import (
	"context"
	"shop-service/internal/domain"

	"github.com/gofiber/fiber/v2"

	authUsecase "shop-service/internal/usecase/auth"
)

type (
	AuthUsecase interface {
		Register(ctx context.Context, req authUsecase.RegisterRequest) (*domain.Session, error)
		Login(ctx context.Context, req authUsecase.LoginRequest) (*domain.Session, error)
		CheckSession(ctx context.Context, session *domain.Session) error
	}
)

type AuthHandlers struct {
	authUC AuthUsecase
}

func NewAuthHandlers(authUC AuthUsecase) *AuthHandlers {
	return &AuthHandlers{
		authUC: authUC,
	}
}

func SetupAuthRoutes(app *fiber.App, handlers *AuthHandlers) {
	app.Post("/auth/register", handlers.Register)
	app.Post("/auth/login", handlers.Login)
}
