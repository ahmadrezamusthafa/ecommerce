package entity

import "time"

type Cart struct {
	ID        string     `db:"id" json:"id,omitempty"`
	UserID    string     `db:"user_id" json:"user_id,omitempty"`
	CreatedAt *time.Time `db:"created_at" json:"-"`
	UpdatedAt *time.Time `db:"updated_at" json:"-"`
	Items     []CartItem `json:"cart_items,omitempty"`
}
