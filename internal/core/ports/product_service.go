package ports

import (
	"context"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
)

type IProductService interface {
	GetAllProducts(ctx context.Context) ([]entity.Product, error)
	GetProductByID(ctx context.Context, productID int) (entity.Product, error)
	SearchProducts(ctx context.Context, query string) ([]entity.Product, error)
}
