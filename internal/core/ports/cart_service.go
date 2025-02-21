package ports

import (
	"context"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
)

type ICartService interface {
	GetCart(ctx context.Context, userID int) (entity.Cart, error)
	AddItemToCart(ctx context.Context, userID int, item entity.CartItem) error
	RemoveItemFromCart(ctx context.Context, userID int, productID int) error
}
