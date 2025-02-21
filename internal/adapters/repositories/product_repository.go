package repositories

import (
	"context"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/ports"
	"github.com/ahmadrezamusthafa/ecommerce/internal/shared/constants"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ports.IProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) GetAll(ctx context.Context) ([]entity.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultHTTPReadTimeout)
	defer cancel()

	var products []entity.Product
	err := r.db.WithContext(ctx).Find(&products).Error
	return products, err
}

func (r *productRepository) GetByID(ctx context.Context, productID string) (entity.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultHTTPReadTimeout)
	defer cancel()

	var product entity.Product
	err := r.db.WithContext(ctx).Where("id = ?", productID).First(&product).Error
	return product, err
}

func (r *productRepository) Search(ctx context.Context, query string) ([]entity.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultHTTPReadTimeout)
	defer cancel()

	var products []entity.Product
	searchQuery := "%" + query + "%"
	err := r.db.Where("name ILIKE ? OR description ILIKE ?", searchQuery, searchQuery).
		Select("id, name, description, price").
		Find(&products).Error

	return products, err
}
