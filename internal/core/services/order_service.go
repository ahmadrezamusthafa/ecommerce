package services

import (
	"context"
	"fmt"
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
	infraContainer  *InfraContainer
}

func NewOrderService(sessionCfg *session.Config,
	infraContainer *InfraContainer,
	orderRepository ports.IOrderRepository,
	cartService ports.ICartService) ports.IOrderService {
	return &orderService{
		sessionCfg:      sessionCfg,
		infraContainer:  infraContainer,
		orderRepository: orderRepository,
		cartService:     cartService,
	}
}

func (s *orderService) CreateOrder(ctx context.Context, userID int) ([]entity.Order, error) {
	var err error
	tx := s.infraContainer.DB.Begin()
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", tx.Error)
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			logger.Errorf("Panic exception: %v", r)
		} else if err != nil {
			tx.Rollback()
			logger.Errorf("Unhandled exception: %v", r)
		}
	}()

	cart, err := s.cartService.GetCart(ctx, userID)
	if err != nil {
		return nil, err
	}

	orders, err := s.prepareOrdersFromCart(cart, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare orders: %w", err)
	}

	orders, err = s.orderRepository.CreateOrder(ctx, tx, orders)
	if err != nil {
		return nil, fmt.Errorf("failed to create orders: %w", err)
	}

	err = s.cartService.RemoveAllItemsFromCart(ctx, tx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to remove cart items: %w", err)
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return orders, nil
}

func (s *orderService) GetTopCustomers(ctx context.Context, limit int) ([]entity.CustomerTransaction, error) {
	orders, err := s.orderRepository.GetTopCustomers(ctx, limit)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *orderService) prepareOrdersFromCart(cart entity.Cart, userID int) ([]entity.Order, error) {
	var orders []entity.Order
	for _, cartItem := range cart.Items {
		if cartItem.Product == nil {
			logger.Infof("Skipped product ID %v due to missing price", cartItem.ProductID)
			continue
		}

		quantity := cartItem.Quantity
		price := cartItem.Product.Price
		amount := float64(quantity) * price

		order := entity.Order{
			CustomerID: userID,
			ProductID:  cartItem.ProductID,
			Quantity:   quantity,
			Amount:     amount,
			OrderDate:  time.Now(),
		}
		orders = append(orders, order)
	}

	if len(orders) == 0 {
		return nil, fmt.Errorf("no valid products found in the cart")
	}

	return orders, nil
}
