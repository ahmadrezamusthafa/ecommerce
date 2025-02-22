package ports

import (
	"context"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
)

type IOrderRepository interface {
	CreateOrder(ctx context.Context, orders []entity.Order) ([]entity.Order, error)
	GetTopCustomers(ctx context.Context, limit int) ([]entity.CustomerTransaction, error)
}
