package db

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"light-ms/pkg/common/config"
)

const (
	dsnFormat = "postgres://%s:%s@%s:%s/%s"
)

type Conn struct {
	pool *pgxpool.Pool
}

func NewConn(ctx context.Context, conf config.Dsn) (Conn, error) {
	cfg, err := pgxpool.ParseConfig(dsnToStr(conf))
	if err != nil {
		return Conn{}, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return Conn{}, err
	}

	if err = pool.Ping(ctx); err != nil {
		return Conn{}, err
	}

	return Conn{pool: pool}, nil
}

func (c Conn) Close() {
	c.pool.Close()
}

func (c Conn) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return c.pool.Query(ctx, sql, args...)
}

func (c Conn) Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	return c.pool.Exec(ctx, sql, args...)
}

func QueryOne[T any](ctx context.Context, cn Conn, q string, args ...any) (T, error) {
	var resp T

	if err := pgxscan.Get(ctx, cn, &resp, q, args...); err != nil {
		return resp, err
	}

	return resp, nil
}

func dsnToStr(c config.Dsn) string {
	return fmt.Sprintf(dsnFormat, c.User, c.Password, c.Host, c.Port, c.Database)
}
