package main

import (
	"7kzu-order-service/internal/config"
	"7kzu-order-service/internal/server"
	"7kzu-order-service/pkg/logger"
	"github.com/joho/godotenv"
)

func main() {
	logger.New()
	err := godotenv.Load(".env")

	if err != nil {
		logger.Fatal("Error  .env file")
	}

	cfg, err := config.InitConfig()
	if err != nil {
		logger.Fatal(err)
	}
	s := server.New(cfg)
	logger.Fatal(s.Run())
}
