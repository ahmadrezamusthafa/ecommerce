package entity

import "time"

type Order struct {
	ID         int        `db:"id" json:"id"`
	CustomerID int        `db:"customer_id" json:"customer_id"`
	ProductID  int        `db:"product_id" json:"product_id"`
	Amount     float64    `db:"amount" json:"amount"`
	OrderDate  *time.Time `db:"order_date" json:"order_date"`
}
