package entity

import "time"

type CartItem struct {
	ID        string     `db:"id" json:"id"`
	CartID    string     `db:"cart_id" json:"cart_id"`
	ProductID string     `db:"product_id" json:"product_id"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}
