package requests

type CreateOrderRequest struct {
	UserId int32 `json:"user_id"`
	Amount int64 `json:"amount"`
}
