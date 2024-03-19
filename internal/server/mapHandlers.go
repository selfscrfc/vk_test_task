package server

import (
	"log/slog"
	"vk_test_task/config"
	api_delivery "vk_test_task/internal/api/delivery"
	api_repository "vk_test_task/internal/api/repository/postgres"
	"vk_test_task/internal/api/repository/redis"
	api_usecase "vk_test_task/internal/api/usecase"
	"vk_test_task/internal/server/delivery/mapRoutes"
)

func MapHandlers(cfg *config.Config, logger *slog.Logger) {
	apiRepo := api_repository.NewRepository(cfg, logger)

	redisRepo := redis.New(cfg, logger)

	apiUc := api_usecase.New(cfg, logger, apiRepo, redisRepo)

	apiHandler := api_delivery.New(cfg, logger, apiUc)

	mapRoutes.MapApiRoutes(cfg, logger, apiHandler)
}
