package requests

type CreateOrderRequest struct {
	UserId int32 `json:"user_id"`
	Amount int64 `json:"amount"`
}

type UpdateStatusRequest struct {
	OrderId int32  `json:"order_id"`
	Status  string `json:"status"`
}
