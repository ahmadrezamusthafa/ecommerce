package requests

type WithdrawBalanceRequest struct {
	Amount float64 `json:"amount" binding:"required"`
}
