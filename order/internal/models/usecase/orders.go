package usecase

type CreateOrderDto struct {
	UserId int32
	Amount int64
}

type UpdateStatusDto struct {
	OrderId int32
	Status  string
}
