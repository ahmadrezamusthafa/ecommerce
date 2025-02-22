package services

import (
	"context"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/session"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/ports"
	"github.com/ahmadrezamusthafa/ecommerce/pkg/logger"
	"time"
)

type orderService struct {
	orderRepository ports.IOrderRepository
	cartService     ports.ICartService
	sessionCfg      *session.Config
}

func NewOrderService(sessionCfg *session.Config, orderRepository ports.IOrderRepository,
	cartService ports.ICartService) ports.IOrderService {
	return &orderService{
		sessionCfg:      sessionCfg,
		orderRepository: orderRepository,
		cartService:     cartService,
	}
}

func (s *orderService) CreateOrder(ctx context.Context, userID int) ([]entity.Order, error) {
	cart, err := s.cartService.GetCart(ctx, userID)
	if err != nil {
		return []entity.Order{}, err
	}

	var orders []entity.Order
	for _, cartItem := range cart.Items {
		if cartItem.Product == nil {
			logger.Infof("Skipped product ID %v caused by price not found", cartItem.ProductID)
			continue
		}

		quantity := cartItem.Quantity
		price := cartItem.Product.Price
		amount := float64(quantity) * price

		order := entity.Order{
			CustomerID: userID,
			ProductID:  cartItem.ProductID,
			Quantity:   cartItem.Quantity,
			Amount:     amount,
			OrderDate:  time.Now(),
		}
		orders = append(orders, order)
	}

	return s.orderRepository.CreateOrder(ctx, orders)
}

func (s *orderService) GetTopCustomers(ctx context.Context, limit int) ([]entity.CustomerTransaction, error) {
	orders, err := s.orderRepository.GetTopCustomers(ctx, limit)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
