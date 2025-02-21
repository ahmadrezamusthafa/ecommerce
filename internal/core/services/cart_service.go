package services

import (
	"context"
	"errors"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/session"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/ports"
	"gorm.io/gorm"
)

type cartService struct {
	cartRepository ports.ICartRepository
	sessionCfg     *session.Config
}

func NewCartService(sessionCfg *session.Config, productRepository ports.ICartRepository) ports.ICartService {
	return &cartService{
		sessionCfg:     sessionCfg,
		cartRepository: productRepository,
	}
}

func (s *cartService) GetCart(ctx context.Context, userID int) (entity.Cart, error) {
	cart, err := s.cartRepository.GetCartByUserID(ctx, userID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.Cart{}, nil
	}
	return cart, err
}

func (s *cartService) AddItemToCart(ctx context.Context, userID int, item entity.CartItem) error {
	cart, err := s.cartRepository.GetCartByUserID(ctx, userID)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		cart, err = s.cartRepository.CreateCart(ctx, entity.Cart{UserID: userID})
		if err != nil {
			return err
		}
	}
	return s.cartRepository.AddItemToCart(ctx, cart.ID, item)
}

func (s *cartService) RemoveItemFromCart(ctx context.Context, userID int, productID int) error {
	cart, err := s.cartRepository.GetCartByUserID(ctx, userID)
	if err != nil {
		return err
	}
	return s.cartRepository.RemoveItemFromCart(ctx, cart.ID, productID)
}
