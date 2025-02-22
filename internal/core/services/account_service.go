package services

import (
	"context"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/session"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/ports"
	"gorm.io/gorm"
)

type accountService struct {
	accountRepository ports.IAccountRepository
	sessionCfg        *session.Config
}

func NewAccountService(sessionCfg *session.Config, accountRepository ports.IAccountRepository) ports.IAccountService {
	return &accountService{
		sessionCfg:        sessionCfg,
		accountRepository: accountRepository,
	}
}

func (s *accountService) GetAccountByUserID(ctx context.Context, userID int) (entity.Account, error) {
	return s.accountRepository.GetAccountByUserID(ctx, userID)
}

func (s *accountService) UpdateAccountBalance(ctx context.Context, tx *gorm.DB, userID int, balance float64) error {
	return s.accountRepository.UpdateAccountBalance(ctx, tx, userID, balance)
}
