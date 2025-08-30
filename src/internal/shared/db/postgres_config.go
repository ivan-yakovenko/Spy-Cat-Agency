package db

import (
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDb(ctx context.Context) (*pgxpool.Pool, error) {

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")
	host := os.Getenv("POSTGRES_HOST")

	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)

	config, err := pgxpool.ParseConfig(dataSourceName)

	if err != nil {
		return nil, error_handler.ErrorHandler(err, "Error parsing pgx config")
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)

	if err != nil {
		return nil, error_handler.ErrorHandler(err, err.Error())
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, error_handler.ErrorHandler(err, err.Error())
	}

	return pool, nil

}
