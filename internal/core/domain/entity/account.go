package entity

import "time"

type Account struct {
	ID        int        `db:"id" json:"id"`
	UserID    int        `db:"user_id" json:"user_id"`
	Balance   float64    `db:"balance" json:"balance"`
	CreatedAt *time.Time `db:"created_at" json:"-"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at,omitempty"`
}
