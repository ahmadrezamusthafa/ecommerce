package ports

import (
	"context"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
	"gorm.io/gorm"
)

type IOrderRepository interface {
	CreateOrder(ctx context.Context, tx *gorm.DB, orders []entity.Order) ([]entity.Order, error)
	GetTopCustomers(ctx context.Context, limit int) ([]entity.CustomerTransaction, error)
	GetByUserID(ctx context.Context, userID int) ([]entity.Order, error)
}
