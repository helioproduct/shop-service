// main.go
package main

import (
	"log"
	"shop-service/config"
	"shop-service/internal/server"
	"shop-service/pkg/logger"
)

func main() {
	logger.InitLogger()
	cfg := config.MustLoadConfig()

	srv := server.NewServer(cfg)
	if err := srv.Initialize(); err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}

	if err := srv.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
