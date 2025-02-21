package ports

import (
	"context"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
)

type ICartRepository interface {
	CreateCart(ctx context.Context, cart entity.Cart) (entity.Cart, error)
	GetCartByUserID(ctx context.Context, userID string) (entity.Cart, error)
	AddItemToCart(ctx context.Context, cartID string, item entity.CartItem) error
	RemoveItemFromCart(ctx context.Context, cartID string, productID string) error
}
