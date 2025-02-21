package entity

import "time"

type CartItem struct {
	ID        string     `db:"id" json:"id,omitempty"`
	CartID    string     `db:"cart_id" json:"cart_id,omitempty"`
	ProductID string     `db:"product_id" json:"product_id,omitempty"`
	Quantity  int        `db:"quantity" json:"quantity"`
	CreatedAt *time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at,omitempty"`
}
