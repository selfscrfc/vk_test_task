package api_usecase

import (
	"log/slog"
	"vk_test_task/config"
	"vk_test_task/internal/api"
)

type UseCase struct {
	cfg    *config.Config
	logger *slog.Logger
	db     api.RepositoryInterface
	rdb    api.TokenRepositoryInterface
}

func New(cfg *config.Config, logger *slog.Logger, db api.RepositoryInterface, tokenRepo api.TokenRepositoryInterface) UseCase {
	return UseCase{
		cfg:    cfg,
		logger: logger,
		db:     db,
		rdb:    tokenRepo,
	}
}
