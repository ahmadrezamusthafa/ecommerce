package repositories

import (
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/ports"
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
