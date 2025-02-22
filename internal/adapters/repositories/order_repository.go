package repositories

import (
	"context"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/ports"
	"github.com/ahmadrezamusthafa/ecommerce/internal/shared/constants"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) ports.IOrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (r *orderRepository) CreateOrder(ctx context.Context, tx *gorm.DB, orders []entity.Order) ([]entity.Order, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultHTTPWriteTimeout)
	defer cancel()

	if err := tx.WithContext(ctx).Create(&orders).Error; err != nil {
		return []entity.Order{}, err
	}
	return orders, nil
}

func (r *orderRepository) GetTopCustomers(ctx context.Context, limit int) ([]entity.CustomerTransaction, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultHTTPReadTimeout)
	defer cancel()

	var orders []entity.CustomerTransaction
	err := r.db.WithContext(ctx).Table("orders").
		Select("orders.customer_id, SUM(orders.amount) AS total_amount, users.name AS customer_name, users.email AS customer_email").
		Joins("LEFT JOIN users ON orders.customer_id = users.id").
		Where("order_date >= NOW() - INTERVAL '1 MONTH'").
		Group("customer_id, customer_name, customer_email").
		Order("total_amount DESC").
		Limit(limit).
		Find(&orders).Error

	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *orderRepository) GetByUserID(ctx context.Context, userID int) ([]entity.Order, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultHTTPReadTimeout)
	defer cancel()

	var orders []entity.Order
	err := r.db.WithContext(ctx).Where("customer_id = ?", userID).Find(&orders).Error
	return orders, err
}
