package usecase

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"light-ms/order/internal/models/entities"
	"light-ms/order/internal/models/usecase"
)

type OrdersRepository interface {
	CreateOrder(ctx context.Context, order entities.Order) error
	GetById(ctx context.Context, id pgtype.Int4) (entities.Order, error)
	UpdateStatus(ctx context.Context, id pgtype.Int4, status pgtype.Text) error
}

type OrdersUcase struct {
	repo OrdersRepository
}

func NewOrdersUcase(repo OrdersRepository) OrdersUcase {
	return OrdersUcase{repo: repo}
}

func (u OrdersUcase) CreateOrder(ctx context.Context, dto usecase.CreateOrderDto) error {
	order := entities.Order{
		UserId: pgtype.Int4{Int32: dto.UserId, Valid: true},
		Amount: pgtype.Int8{Int64: dto.Amount, Valid: true},
	}

	if err := u.repo.CreateOrder(ctx, order); err != nil {
		return fmt.Errorf("db: %v", err)
	}

	return nil
}

func (u OrdersUcase) GetById(ctx context.Context, id int32) (entities.Order, error) {
	order, err := u.repo.GetById(ctx, pgtype.Int4{Int32: id, Valid: true})
	if err != nil {
		return order, fmt.Errorf("db: %v", err)
	}

	return order, nil
}

func (u OrdersUcase) UpdateStatus(ctx context.Context, dto usecase.UpdateStatusDto) error {
	id := pgtype.Int4{Int32: dto.OrderId, Valid: true}
	status := pgtype.Text{String: dto.Status, Valid: true}

	if err := u.repo.UpdateStatus(ctx, id, status); err != nil {
		return fmt.Errorf("db: %v", err)
	}

	return nil
}
