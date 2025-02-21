package repositories

import (
	"context"
	"errors"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/ports"
	"github.com/ahmadrezamusthafa/ecommerce/internal/shared/constants"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) ports.ICartRepository {
	return &cartRepository{
		db: db,
	}
}

func (r *cartRepository) CreateCart(ctx context.Context, cart entity.Cart) (entity.Cart, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultHTTPWriteTimeout)
	defer cancel()

	cart.ID = uuid.NewString()
	if err := r.db.WithContext(ctx).Create(&cart).Error; err != nil {
		return entity.Cart{}, err
	}
	return cart, nil
}

func (r *cartRepository) GetCartByUserID(ctx context.Context, userID string) (entity.Cart, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultHTTPReadTimeout)
	defer cancel()

	var cart entity.Cart
	err := r.db.WithContext(ctx).Preload("Items.Product").Where("user_id = ?", userID).First(&cart).Error
	if err != nil {
		return entity.Cart{}, err
	}
	return cart, nil
}

func (r *cartRepository) AddItemToCart(ctx context.Context, cartID string, item entity.CartItem) error {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultHTTPWriteTimeout)
	defer cancel()

	item.ID = uuid.NewString()
	item.CartID = cartID
	return r.db.WithContext(ctx).Create(&item).Error
}

func (r *cartRepository) RemoveItemFromCart(ctx context.Context, cartID string, productID string) error {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultHTTPWriteTimeout)
	defer cancel()

	result := r.db.WithContext(ctx).Where("cart_id = ? AND product_id = ?", cartID, productID).Delete(&entity.CartItem{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("item not found in cart")
	}
	return nil
}
