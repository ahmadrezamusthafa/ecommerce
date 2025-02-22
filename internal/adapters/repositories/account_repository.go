package repositories

import (
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/ports"
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
