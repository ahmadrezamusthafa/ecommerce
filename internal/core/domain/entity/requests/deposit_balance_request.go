package requests

type DepositBalanceRequest struct {
	Amount float64 `json:"amount" binding:"required"`
}
