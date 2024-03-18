package postgres

import (
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"vk_test_task/config"
	log "vk_test_task/pkg/logger"
)

type Repository struct {
	cfg    *config.Config
	db     *sqlx.DB
	logger *slog.Logger
}

func NewRepository(cfg *config.Config, logger *slog.Logger) Repository {
	db, err := sqlx.Connect("pgx", fmt.Sprintf("host=%s port=%s user=%s password=%s database=%s sslmode=%s",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
		cfg.Postgres.SSLMode,
	))
	if err != nil {
		log.Fatalf(logger, "postgres connection error: %s", err.Error())
	}
	logger.Debug("postgres database connected")

	return Repository{
		db:  db,
		cfg: cfg,
	}
}
