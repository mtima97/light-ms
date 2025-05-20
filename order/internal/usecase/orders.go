package usecase

import "light-ms/order/internal/repository"

type OrdersUcase struct {
	repo repository.Db
}

func NewOrdersUcase(repo repository.Db) OrdersUcase {
	return OrdersUcase{repo: repo}
}
