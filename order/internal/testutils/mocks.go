package testutils

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"light-ms/order/internal/models/entities"
)

type MockOrdersRepository struct {
}

func (m MockOrdersRepository) CreateOrder(ctx context.Context, order entities.Order) error {
	return nil
}

func (m MockOrdersRepository) GetById(ctx context.Context, id pgtype.Int4) (entities.Order, error) {
	return entities.Order{}, nil
}

func (m MockOrdersRepository) Get(ctx context.Context) ([]entities.Order, error) {
	return []entities.Order{}, nil
}

func (m MockOrdersRepository) UpdateStatus(ctx context.Context, id pgtype.Int4, status pgtype.Text) error {
	return nil
}
