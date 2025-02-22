package ports

import (
	"context"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
	"gorm.io/gorm"
)

type ICartRepository interface {
	CreateCart(ctx context.Context, cart entity.Cart) (entity.Cart, error)
	GetCartByUserID(ctx context.Context, userID int) (entity.Cart, error)
	AddItemToCart(ctx context.Context, cartID int, item entity.CartItem) error
	RemoveItemFromCart(ctx context.Context, cartID int, productID int) error
	RemoveAllItemsFromCart(ctx context.Context, tx *gorm.DB, cartID int) error
	GetItemByProductID(ctx context.Context, cartID, productID int) (entity.CartItem, error)
	UpdateCartItem(ctx context.Context, item entity.CartItem) error
}
