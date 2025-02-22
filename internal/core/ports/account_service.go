package ports

import (
	"context"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
	"gorm.io/gorm"
)

type IAccountService interface {
	GetAccountByUserID(ctx context.Context, userID int) (entity.Account, error)
	UpdateAccountBalance(ctx context.Context, tx *gorm.DB, userID int, balance float64) error
}
