package middleware

import (
	"context"
	"errors"
	"shop-service/internal/domain"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type AuthUsecase interface {
	CheckSession(ctx context.Context, session *domain.Session) error
}

type AuthMiddleware struct {
	authUsecase AuthUsecase
}

func NewAuthMiddleware(authUsecase AuthUsecase) *AuthMiddleware {
	return &AuthMiddleware{authUsecase: authUsecase}
}

// Middleware для проверки JWT-токена
func (m *AuthMiddleware) Handle() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Извлекаем токен из заголовка Authorization
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing authorization header"})
		}

		// Ожидаем формат: Bearer <token>
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid authorization header format"})
		}

		token := parts[1]

		// Создаем сессию из токена
		session := &domain.Session{
			Token: token,
		}

		if err := m.authUsecase.CheckSession(c.Context(), session); err != nil {
			if errors.Is(err, domain.ErrInvalidToken) {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
			}

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "internal server error"})
		}

		// Добавляем сессию в контекст
		c.Locals("session", session)

		// Переходим к следующему обработчику
		return c.Next()
	}
}

// Получить сессию из контекста
func GetSessionFromContext(c *fiber.Ctx) (*domain.Session, bool) {
	session, ok := c.Locals("session").(*domain.Session)
	return session, ok
}
