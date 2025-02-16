package main

import (
	"database/sql"
	"fmt"
	"log"
	"shop-service/config"
	authHandlers "shop-service/internal/handler/auth"
	userRepository "shop-service/internal/repository/user"
	"shop-service/internal/usecase/auth"

	trmsql "github.com/avito-tech/go-transaction-manager/drivers/sql/v2"
	"github.com/gofiber/fiber/v2"

	_ "github.com/lib/pq"
	// _ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	cfg := config.MustLoadConfig()
	fmt.Println("Postgres DSN:", cfg.PostgresConfig.DSN())

	db, err := sql.Open("postgres", cfg.PostgresConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}
	fmt.Println("Successfully connected to the database.")

	txGetter := trmsql.DefaultCtxGetter
	// trManager := manager.Must(trmsql.NewDefaultFactory(db))

	userRepo := userRepository.NewUserRepository(db, txGetter)

	authUC := auth.NewAuthUsecase(cfg, userRepo)

	authHandler := authHandlers.NewAuthHandlers(authUC)
	app := fiber.New()

	authHandlers.SetupAuthRoutes(app, authHandler)

	log.Printf("Starting server on port %s...", cfg.ServerConfig.Port)
	if err := app.Listen(":" + string(cfg.ServerConfig.Port)); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

// func main() {

// 	// 5. Инициализируем хендлеры
// 	authHandlers := handlers.NewAuthHandlers(authUC)

// 	// 6. Создаем Fiber приложение
// 	app := fiber.New()

// 	// 7. Настраиваем маршруты
// 	handlers.SetupAuthRoutes(app, authHandlers)

// }
