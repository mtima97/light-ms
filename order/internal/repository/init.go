package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

func init() {
	dsn := "postgres://DB_USER:DB_PASSWORD@localhost:CUSTOM_DB_PORT/postgres"

	pgxpool.New(context.Background(), dsn)
}
