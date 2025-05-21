package usecase

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"light-ms/order/internal/models/entities"
	"light-ms/order/internal/models/usecase"
	"light-ms/order/internal/repository"
)

type OrdersUcase struct {
	repo repository.Db
}

func NewOrdersUcase(repo repository.Db) OrdersUcase {
	return OrdersUcase{repo: repo}
}

func (u OrdersUcase) CreateOrder(ctx context.Context, dto usecase.CreateOrderDto) error {
	order := entities.Order{
		UserId: pgtype.Int4{Int32: dto.UserId, Valid: true},
		Amount: pgtype.Int8{Int64: dto.Amount, Valid: true},
	}

	if err := u.repo.CreateOrder(ctx, order); err != nil {
		return fmt.Errorf("create order in db: %v", err)
	}

	return nil
}
