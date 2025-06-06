package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"light-ms/order/internal/models/entities"
	cmnconf "light-ms/pkg/common/config"
	"light-ms/pkg/common/db"
	"log"
)

type OrderRepo struct {
	conn db.Conn
}

func NewOrderRepository(ctx context.Context, cfg cmnconf.Dsn) OrderRepo {
	conn, err := db.NewConn(ctx, cfg)
	if err != nil {
		log.Fatalf("db: %v", err)
	}

	return OrderRepo{conn: conn}
}

func (r OrderRepo) CreateOrder(ctx context.Context, order entities.Order) error {
	sql := "insert into orders.orders (user_id, amount) values ($1, $2);"

	_, err := r.conn.Exec(ctx, sql, order.UserId, order.Amount)
	if err != nil {
		return err
	}

	return nil
}

func (r OrderRepo) GetById(ctx context.Context, id pgtype.Int4) (entities.Order, error) {
	sql := "select id, user_id, amount, status, created_at, updated_at from orders.orders where id = $1;"

	order, err := db.QueryOne[entities.Order](ctx, r.conn, sql, id)
	if err != nil {
		return entities.Order{}, err
	}

	return order, nil
}

func (r OrderRepo) UpdateStatus(ctx context.Context, id pgtype.Int4, status pgtype.Text) error {
	sql := "update orders.orders set status = $1 where id = $2;"

	_, err := r.conn.Exec(ctx, sql, status, id)
	if err != nil {
		return err
	}

	return nil
}

func (r OrderRepo) Close() {
	r.conn.Close()
}
