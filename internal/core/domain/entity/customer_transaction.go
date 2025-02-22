package entity

type CustomerTransaction struct {
	CustomerID    int     `json:"customer_id"`
	CustomerName  string  `json:"customer_name"`
	CustomerEmail string  `json:"customer_email"`
	TotalAmount   float64 `json:"total_amount"`
}
