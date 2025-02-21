package entity

type CartItem struct {
	ID        int      `db:"id" json:"id,omitempty"`
	CartID    int      `db:"cart_id" json:"cart_id,omitempty"`
	ProductID int      `db:"product_id" json:"product_id,omitempty"`
	Quantity  int      `db:"quantity" json:"quantity"`
	Product   *Product `json:"product,omitempty"`
}
