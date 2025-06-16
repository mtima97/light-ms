package requests

type CreateOrderRequest struct {
	Amount int64 `json:"amount" binding:"required,gte=1"`
}

type UpdateStatusRequest struct {
	OrderId int32  `json:"order_id" binding:"required,gte=1"`
	Status  string `json:"status" binding:"required"`
}
