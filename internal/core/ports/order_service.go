package ports

import (
	"context"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
)

type IOrderService interface {
	CreateOrder(ctx context.Context, userID int) ([]entity.Order, error)
	GetTopCustomers(ctx context.Context, limit int) ([]entity.CustomerTransaction, error)
}
