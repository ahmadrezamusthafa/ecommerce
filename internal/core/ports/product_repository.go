package ports

import (
	"context"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
)

type IProductRepository interface {
	GetAll(ctx context.Context) ([]entity.Product, error)
	GetByID(ctx context.Context, productID int) (entity.Product, error)
	Search(ctx context.Context, query string) ([]entity.Product, error)
}
