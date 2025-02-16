package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"shop-service/config"
	transferRepository "shop-service/internal/repository/transfer"
	userRepository "shop-service/internal/repository/user"
	transferUsecase "shop-service/internal/usecase/transfer"
	"time"

	trmsql "github.com/avito-tech/go-transaction-manager/drivers/sql/v2"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"

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
	trManager := manager.Must(trmsql.NewDefaultFactory(db))

	userRepo := userRepository.NewUserRepository(db, txGetter)
	transferRepo := transferRepository.NewTransferRepository(db, txGetter)

	transferUC := transferUsecase.NewTransferUsecase(trManager, transferRepo, userRepo)

	// req := transferUsecase.SendCoinsRequest{
	// From:   "carol",
	// To:     "alice",
	// Amount: 1000,
	// }

	time.Sleep(3 * time.Second)
	start := time.Now()

	coins, err := transferUC.GetReceivedCoinsHistory(context.Background(), "alice")
	// err = transferUC.SendCoins(context.Background(), req)

	if err != nil {
		log.Fatalf("failed to send coins: %v", err)
	}
	fmt.Println(time.Since(start))

	fmt.Println(coins)

}
