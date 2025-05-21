package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"light-ms/order/internal/config"
	"log"
)

type Db struct {
	pool *pgxpool.Pool
}

func NewDbConn(ctx context.Context, glConf config.Config) Db {
	conf, err := pgxpool.ParseConfig(getDsn(glConf))
	if err != nil {
		log.Fatalf("db: %v", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, conf)
	if err != nil {
		log.Fatalf("db: %v", err)
	}

	if err = pool.Ping(ctx); err != nil {
		log.Fatalf("db: %v", err)
	}

	return Db{pool: pool}
}

func (d Db) Close() {
	d.pool.Close()
}

func getDsn(conf config.Config) string {
	f := "postgres://%s:%s@%s:%s/%s"

	return fmt.Sprintf(f, conf.DbUser, conf.DbPassword, conf.DbHost, conf.DbPort, conf.DbName)
}
