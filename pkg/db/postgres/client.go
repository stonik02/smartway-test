package postgres

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v3/log"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"test-task/internal/config"
)

func NewPgConnection(cfg *config.DB) (*pgxpool.Pool, error) {
	dsn := getDSNFromConfig(cfg)
	pgxConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		panic(err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), pgxConfig)
	if err != nil {
		log.Errorf("Unable to connect pool to database: %v\n", err)
		os.Exit(1)
	}

	return pool, nil
}

func getDSNFromConfig(cfg *config.DB) string {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database,
	)

	return url
}
