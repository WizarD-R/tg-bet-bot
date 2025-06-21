package db

import (
	"RushBananaBet/pkg/logger"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(ctx context.Context, dbURL string) *pgxpool.Pool {
	// postgres://user:pass@localhost:5432/dbname?sslmode=disable

	//c := pgxpool.New()

	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		logger.Fatal("Error parse config for pgxpool", "", "", "db - NewPostgresPool()", err)
		return nil
	}

	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnIdleTime = 5 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return fmt.Errorf("error creating pool: %w", err)
	}

	// Проверим соединение
	if err := pool.Ping(ctx); err != nil {
		return fmt.Errorf("could not ping database: %w", err)
	}
}
