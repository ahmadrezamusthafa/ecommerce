package repositories

import (
	"context"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/ports"
	"github.com/ahmadrezamusthafa/ecommerce/internal/shared/constants"
	"gorm.io/gorm"
)

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) ports.IAccountRepository {
	return &accountRepository{
		db: db,
	}
}

func (r *accountRepository) GetAccountByUserID(ctx context.Context, userID int) (entity.Account, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultHTTPReadTimeout)
	defer cancel()

	var account entity.Account
	err := r.db.WithContext(ctx).FirstOrCreate(&account, entity.Account{UserID: userID}).Error
	if err != nil {
		return entity.Account{}, err
	}
	return account, nil
}

func (r *accountRepository) IncreaseAccountBalance(ctx context.Context, tx *gorm.DB, userID int, amount float64) error {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultHTTPWriteTimeout)
	defer cancel()

	if tx == nil {
		tx = r.db
	}

	return tx.WithContext(ctx).
		Model(&entity.Account{}).
		Where("user_id = ?", userID).
		Update("balance", gorm.Expr("balance + ?", amount)).
		Error
}

func (r *accountRepository) DecreaseAccountBalance(ctx context.Context, tx *gorm.DB, userID int, amount float64) error {
	ctx, cancel := context.WithTimeout(ctx, constants.DefaultHTTPWriteTimeout)
	defer cancel()

	if tx == nil {
		tx = r.db
	}

	return tx.WithContext(ctx).
		Model(&entity.Account{}).
		Where("user_id = ?", userID).
		Update("balance", gorm.Expr("balance - ?", amount)).
		Error
}
