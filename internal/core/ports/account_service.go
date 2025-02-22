package ports

import (
	"context"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
	"gorm.io/gorm"
)

type IAccountService interface {
	GetAccountByUserID(ctx context.Context, userID int) (entity.Account, error)
	IncreaseAccountBalance(ctx context.Context, tx *gorm.DB, userID int, balance float64) error
	DecreaseAccountBalance(ctx context.Context, tx *gorm.DB, userID int, balance float64) error
	Withdraw(ctx context.Context, userID int, amount float64) error
	Deposit(ctx context.Context, userID int, amount float64) error
}
