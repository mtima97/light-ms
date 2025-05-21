package repository

import (
	"context"
	"light-ms/order/internal/models/entities"
)

func (d Db) CreateOrder(ctx context.Context, order entities.Order) error {
	sql := "insert into orders.orders (user_id, amount) values ($1, $2)"

	_, err := d.pool.Exec(ctx, sql, order.UserId, order.Amount)
	if err != nil {
		return err
	}

	return nil
}
