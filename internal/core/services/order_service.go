package services

import (
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/session"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/ports"
)

type orderService struct {
	orderRepository ports.IOrderRepository
	sessionCfg      *session.Config
}

func NewOrderService(sessionCfg *session.Config, orderRepository ports.IOrderRepository) ports.IOrderService {
	return &orderService{
		sessionCfg:      sessionCfg,
		orderRepository: orderRepository,
	}
}
