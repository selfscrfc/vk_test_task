package main

import (
	"vk_test_task/config"
	_ "vk_test_task/docs"
	"vk_test_task/internal/server"
	tint "vk_test_task/pkg/logger"

	_ "database/sql"
)

// @title VK_TEST_TASK
// @host localhost:9091
// @BasePath /

// @securityDefinitions.apiKey AccessTokenAuth
// @in header
// @name Authorization
func main() {
	cfg := config.ParseConfig()

	logger := tint.NewLogger(cfg.Logger.InFile)

	logger.Info("config and logger successfully started")

	server.Run(cfg, logger)
}
