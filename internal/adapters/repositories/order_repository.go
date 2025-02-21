package repositories

import (
	"context"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/ports"
	"github.com/ahmadrezamusthafa/ecommerce/internal/shared/constants"
	"gorm.io/gorm"
	"time"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) ports.IOrderRepository {
	return &orderRepository{
		db: db,
	}
}
func (r *orderRepository) CreateOrder(ctx context.Context, order entity.Order) (entity.Order, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultHTTPReadTimeout)
	defer cancel()

	order.OrderDate = time.Now()
	if err := r.db.WithContext(ctx).Create(&order).Error; err != nil {
		return entity.Order{}, err
	}
	return order, nil
}

func (r *orderRepository) GetTopCustomers(ctx context.Context, limit int) ([]entity.Order, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultHTTPReadTimeout)
	defer cancel()

	var orders []entity.Order
	err := r.db.WithContext(ctx).
		Select("customer_id, SUM(amount) AS total_amount").
		Where("order_date >= NOW() - INTERVAL '1 MONTH'").
		Group("customer_id").
		Order("total_amount DESC").
		Limit(limit).
		Find(&orders).Error

	if err != nil {
		return nil, err
	}
	return orders, nil
}
