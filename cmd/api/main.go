package main

import (
	"vk_test_task/config"
	"vk_test_task/internal/server"
	tint "vk_test_task/pkg/logger"
)

func main() {
	cfg := config.ParseConfig()

	logger := tint.NewLogger(cfg.Logger.InFile)

	logger.Info("config and logger successfully started")

	server.Run(cfg, logger)
}
