package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"vk_test_task/config"
	logger2 "vk_test_task/pkg/logger"
)

func Run(cfg *config.Config, logger *slog.Logger) {
	MapHandlers(cfg, logger)

	logger.Info("server successfully started")

	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Server.Port), nil); err != nil {
		logger2.Fatalf(logger, "server run fatal error: %s", err.Error())
	}
}
