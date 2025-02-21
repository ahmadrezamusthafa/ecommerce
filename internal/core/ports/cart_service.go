package ports

import (
	"context"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
)

type ICartService interface {
	GetCart(ctx context.Context, userID string) (entity.Cart, error)
	AddItemToCart(ctx context.Context, userID string, item entity.CartItem) error
	RemoveItemFromCart(ctx context.Context, userID string, productID string) error
}
