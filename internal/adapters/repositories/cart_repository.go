package repositories

import (
	"context"
	"errors"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/ports"
	"github.com/ahmadrezamusthafa/ecommerce/internal/shared/constants"
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

	if err := r.db.WithContext(ctx).Create(&cart).Error; err != nil {
		return entity.Cart{}, err
	}
	return cart, nil
}

func (r *cartRepository) GetCartByUserID(ctx context.Context, userID int) (entity.Cart, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultHTTPReadTimeout)
	defer cancel()

	var cart entity.Cart
	err := r.db.WithContext(ctx).Preload("Items.Product").Where("user_id = ?", userID).First(&cart).Error
	if err != nil {
		return entity.Cart{}, err
	}
	return cart, nil
}

func (r *cartRepository) AddItemToCart(ctx context.Context, cartID int, item entity.CartItem) error {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultHTTPWriteTimeout)
	defer cancel()

	item.CartID = cartID
	return r.db.WithContext(ctx).Create(&item).Error
}

func (r *cartRepository) RemoveItemFromCart(ctx context.Context, cartID int, productID int) error {
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

func (r *cartRepository) RemoveAllItemsFromCart(ctx context.Context, tx *gorm.DB, cartID int) error {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultHTTPWriteTimeout)
	defer cancel()

	result := tx.WithContext(ctx).Where("cart_id = ?", cartID).Delete(&entity.CartItem{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("item not found in cart")
	}
	return nil
}

func (r *cartRepository) GetItemByProductID(ctx context.Context, cartID, productID int) (entity.CartItem, error) {
	var item entity.CartItem
	err := r.db.WithContext(ctx).
		Where("cart_id = ? AND product_id = ?", cartID, productID).
		First(&item).Error

	return item, err
}

func (r *cartRepository) UpdateCartItem(ctx context.Context, item entity.CartItem) error {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultHTTPWriteTimeout)
	defer cancel()

	return r.db.WithContext(ctx).
		Model(&entity.CartItem{}).
		Where("id = ?", item.ID).
		Updates(map[string]interface{}{
			"quantity": item.Quantity,
		}).Error
}
