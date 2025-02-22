package services

import (
	"context"
	"fmt"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/session"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/ports"
	"github.com/ahmadrezamusthafa/ecommerce/pkg/logger"
	"github.com/go-redis/redis/v8"
	"time"
)

type orderService struct {
	orderRepository ports.IOrderRepository
	cacheRepository ports.ICacheRepository
	cartService     ports.ICartService
	accountService  ports.IAccountService
	sessionCfg      *session.Config
	infraContainer  *InfraContainer
}

func NewOrderService(sessionCfg *session.Config,
	infraContainer *InfraContainer,
	orderRepository ports.IOrderRepository,
	cacheRepository ports.ICacheRepository,
	cartService ports.ICartService,
	accountService ports.IAccountService) ports.IOrderService {
	return &orderService{
		sessionCfg:      sessionCfg,
		infraContainer:  infraContainer,
		orderRepository: orderRepository,
		cacheRepository: cacheRepository,
		cartService:     cartService,
		accountService:  accountService,
	}
}

func (s *orderService) CreateOrder(ctx context.Context, userID int) ([]entity.Order, error) {
	var err error
	dbTx := s.infraContainer.DB.Begin()
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", dbTx.Error)
	}

	defer func() {
		if r := recover(); r != nil {
			dbTx.Rollback()
			logger.Errorf("Panic exception: %v", r)
		} else if err != nil {
			dbTx.Rollback()
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

	totalAmount := 0.0
	for _, order := range orders {
		totalAmount += order.Amount
	}

	err = s.cacheRepository.WatchUserBalance(userID, func(tx *redis.Tx) error {
		balance, err := s.cacheRepository.GetUserBalance(tx, userID)
		if err != nil {
			return err
		}

		if balance <= 0 {
			account, err := s.accountService.GetAccountByUserID(ctx, userID)
			if err != nil {
				return err
			}
			available, err := s.cacheRepository.SetUserBalance(tx, userID, account.Balance)
			if err != nil {
				return err
			}
			if available {
				balance = account.Balance
			} else {
				balance, err = s.cacheRepository.GetUserBalance(tx, userID)
				if err != nil {
					return err
				}
			}
		}

		newBalance := balance - totalAmount
		if newBalance < 0 {
			return fmt.Errorf("insufficient balance: %.f", newBalance)
		}
		err = s.cacheRepository.DecreaseUserBalance(tx, userID, totalAmount)
		if err != nil {
			return err
		}
		err = s.accountService.UpdateAccountBalance(ctx, dbTx, userID, newBalance)
		if err != nil {
			return err
		}
		orders, err = s.orderRepository.CreateOrder(ctx, dbTx, orders)
		if err != nil {
			return fmt.Errorf("failed to create orders: %w", err)
		}
		err = s.cartService.RemoveAllItemsFromCart(ctx, dbTx, userID)
		if err != nil {
			return fmt.Errorf("failed to remove cart items: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	err = dbTx.Commit().Error
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

func (s *orderService) GetByUserID(ctx context.Context, userID int) ([]entity.Order, error) {
	return s.orderRepository.GetByUserID(ctx, userID)
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
