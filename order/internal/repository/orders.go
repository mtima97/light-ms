package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"light-ms/order/internal/models/entities"
)

func (d Db) CreateOrder(ctx context.Context, order entities.Order) error {
	sql := "insert into orders.orders (user_id, amount) values ($1, $2);"

	_, err := d.pool.Exec(ctx, sql, order.UserId, order.Amount)
	if err != nil {
		return err
	}

	return nil
}

func (d Db) GetById(ctx context.Context, id pgtype.Int4) (entities.Order, error) {
	var order entities.Order

	sql := "select id, user_id, amount, status, created_at, updated_at from orders.orders where id = $1;"

	row := d.pool.QueryRow(ctx, sql, id)

	if err := row.Scan(
		&order.Id,
		&order.UserId,
		&order.Amount,
		&order.Status,
		&order.CreatedAt,
		&order.UpdatedAt,
	); err != nil {
		return order, err
	}

	return order, nil
}
