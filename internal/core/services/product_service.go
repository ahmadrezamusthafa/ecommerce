package services

import (
	"context"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/session"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/ports"
)

type productService struct {
	productRepository ports.IProductRepository
	sessionCfg        *session.Config
}

func NewProductService(sessionCfg *session.Config, productRepository ports.IProductRepository) ports.IProductService {
	return &productService{
		sessionCfg:        sessionCfg,
		productRepository: productRepository,
	}
}

func (s *productService) GetAllProducts(ctx context.Context) ([]entity.Product, error) {
	return s.productRepository.GetAll(ctx)
}

func (s *productService) GetProductByID(ctx context.Context, productID string) (entity.Product, error) {
	return s.productRepository.GetByID(ctx, productID)
}

func (s *productService) SearchProducts(ctx context.Context, query string) ([]entity.Product, error) {
	return s.productRepository.Search(ctx, query)
}
