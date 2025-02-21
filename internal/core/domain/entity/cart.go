package entity

type Cart struct {
	ID     int        `db:"id" json:"id,omitempty"`
	UserID int        `db:"user_id" json:"user_id,omitempty"`
	Items  []CartItem `json:"cart_items,omitempty"`
}
