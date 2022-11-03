package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/aibeksarsembayev/onelab/tasks/lab4/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitPostgresDBConn(config *config.Config) (*pgxpool.Pool, error) {
	// 	Example DSN
	// user=jack password=secret host=pg.example.com port=5432 dbname=mydb sslmode=verify-ca pool_max_conns=10

	// Example URL
	// postgres://jack:secret@pg.example.com:5432/mydb?sslmode=verify-ca&pool_max_conns=10

	// Current URL
	// postgres://postgres:postgres@postgres:5432/postgresdb

	URL := fmt.Sprintf("postgres://%s:%s@%s%s/%s", config.Database.DBUser, config.Database.DBPass, config.Database.DBHost, config.Database.DBPort, config.Database.DBName)

	cfg, err := pgxpool.ParseConfig(URL)
	if err != nil {
		return nil, fmt.Errorf("postgres: %w", err)
	}

	cfg.MaxConns = 10 // default 4 or CPU number
	cfg.MaxConnLifetime = 5 * time.Minute

	pool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		return nil, fmt.Errorf("postgres: %w", err)
	}

	// TODO: call function to insert predefined data into database

	return pool, nil
}

// TODO: function to insert predefined data into database
