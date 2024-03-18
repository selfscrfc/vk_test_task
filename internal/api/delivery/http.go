package api_delivery

import (
	"log/slog"
	"vk_test_task/config"
	"vk_test_task/internal/api"
)

type Handler struct {
	cfg    *config.Config
	logger *slog.Logger
	uc     api.UseCaseInterface
}

func New(cfg *config.Config, logger *slog.Logger, uc api.UseCaseInterface) Handler {
	return Handler{
		cfg:    cfg,
		logger: logger,
		uc:     uc,
	}
}
