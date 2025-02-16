package server

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"shop-service/config"
	authHandlers "shop-service/internal/handler/auth"
	infoHandlers "shop-service/internal/handler/info"
	purchaseHandlers "shop-service/internal/handler/purchase"
	transferHanlders "shop-service/internal/handler/transfer"
	middleware "shop-service/internal/middleware"
	productRepository "shop-service/internal/repository/product"
	purchaseRepository "shop-service/internal/repository/purchase"
	transferRepository "shop-service/internal/repository/transfer"
	userRepository "shop-service/internal/repository/user"
	authUsecase "shop-service/internal/usecase/auth"
	purchaseUsecase "shop-service/internal/usecase/purchase"
	transferUsecase "shop-service/internal/usecase/transfer"
	userUsecase "shop-service/internal/usecase/user"
	"shop-service/pkg/logger"

	trmsql "github.com/avito-tech/go-transaction-manager/drivers/sql/v2"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app *fiber.App
	cfg *config.Config
	db  *sql.DB
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Initialize() error {
	db, err := sql.Open("postgres", s.cfg.PostgresConfig.DSN())
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}
	s.db = db

	txGetter := trmsql.DefaultCtxGetter
	trManager := manager.Must(trmsql.NewDefaultFactory(db))

	userRepo := userRepository.NewUserRepository(db, txGetter)
	transferRepo := transferRepository.NewTransferRepository(db, txGetter)
	purchaseRepo := purchaseRepository.NewPurchaseRepository(db, txGetter)
	productRepo := productRepository.NewProductRepository(db, txGetter)

	authUC := authUsecase.NewAuthUsecase(s.cfg, userRepo)
	transferUC := transferUsecase.NewTransferUsecase(trManager, transferRepo, userRepo)
	userUC := userUsecase.NewUserUsecase(userRepo)
	purchaseUC := purchaseUsecase.NewPurchaseUsecase(trManager, purchaseRepo, userRepo, productRepo)

	authHandler := authHandlers.NewAuthHandlers(authUC)
	infoHandler := infoHandlers.NewInfoHandler(purchaseUC, transferUC, userUC)
	purchaseHandler := purchaseHandlers.NewPurchaseHandler(purchaseUC, userUC)
	transferHandler := transferHanlders.NewTransferHandler(transferUC, userUC)

	s.app = fiber.New()
	s.app.Use(middleware.ZerologMiddleware())

	authHandlers.SetupAuthRoutes(s.app, authHandler)
	authMiddleware := middleware.NewAuthMiddleware(authUC)

	auth := s.app.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	api := s.app.Group("/api", authMiddleware.Auth())
	api.Get("/profile", func(c *fiber.Ctx) error {
		session, _ := middleware.GetSessionFromContext(c)
		return c.JSON(fiber.Map{
			"message": "Hello, " + session.Username,
		})
	})
	api.Get("/info", infoHandler.HandleInfo)
	api.Get("/buy/:item", purchaseHandler.HandlePurchase)
	api.Post("/transfer", transferHandler.HandleTransfer)

	return nil
}

func (s *Server) Run() error {
	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		logger.Log.Info().Msg("Starting Fiber server on port 8080...")
		if err := s.app.Listen(":8080"); err != nil {
			logger.Log.Fatal().Err(err).Msg("Failed to start server")
		}
	}()

	<-shutdownChan
	logger.Log.Info().Msg("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.app.ShutdownWithContext(ctx); err != nil {
		logger.Log.Error().Err(err).Msg("Server forced to shutdown")
		return err
	}

	if err := s.db.Close(); err != nil {
		logger.Log.Error().Err(err).Msg("Error closing database connection")
		return err
	}

	logger.Log.Info().Msg("Server stopped gracefully")
	return nil
}
