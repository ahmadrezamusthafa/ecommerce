package services

import (
	"context"
	"errors"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/session"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/ports"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type accountService struct {
	accountRepository ports.IAccountRepository
	cacheRepository   ports.ICacheRepository
	sessionCfg        *session.Config
}

func NewAccountService(sessionCfg *session.Config,
	accountRepository ports.IAccountRepository,
	cacheRepository ports.ICacheRepository) ports.IAccountService {
	return &accountService{
		sessionCfg:        sessionCfg,
		accountRepository: accountRepository,
		cacheRepository:   cacheRepository,
	}
}

func (s *accountService) GetAccountByUserID(ctx context.Context, userID int) (entity.Account, error) {
	return s.accountRepository.GetAccountByUserID(ctx, userID)
}

func (s *accountService) UpdateAccountBalance(ctx context.Context, tx *gorm.DB, userID int, balance float64) error {
	return s.accountRepository.UpdateAccountBalance(ctx, tx, userID, balance)
}

func (s *accountService) Withdraw(ctx context.Context, userID int, amount float64) error {
	err := s.cacheRepository.WatchUserBalance(userID, func(tx *redis.Tx) error {
		balance, err := s.cacheRepository.GetUserBalance(tx, userID)
		if err != nil {
			return err
		}

		if balance <= 0 {
			account, err := s.accountRepository.GetAccountByUserID(ctx, userID)
			if err != nil {
				return err
			}
			available, err := s.cacheRepository.SetUserBalance(tx, userID, account.Balance)
			if err != nil {
				return err
			}
			if available {
				balance = account.Balance
			} else {
				balance, err = s.cacheRepository.GetUserBalance(tx, userID)
				if err != nil {
					return err
				}
			}
		}

		newBalance := balance - amount
		if newBalance < 0 {
			return errors.New("insufficient balance")
		}
		err = s.cacheRepository.DecreaseUserBalance(tx, userID, amount)
		if err != nil {
			return err
		}
		err = s.accountRepository.UpdateAccountBalance(ctx, nil, userID, newBalance)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *accountService) Deposit(ctx context.Context, userID int, amount float64) error {
	err := s.cacheRepository.WatchUserBalance(userID, func(tx *redis.Tx) error {
		balance, err := s.cacheRepository.GetUserBalance(tx, userID)
		if err != nil {
			return err
		}

		if balance <= 0 {
			account, err := s.accountRepository.GetAccountByUserID(ctx, userID)
			if err != nil {
				return err
			}
			available, err := s.cacheRepository.SetUserBalance(tx, userID, account.Balance)
			if err != nil {
				return err
			}
			if available {
				balance = account.Balance
			} else {
				balance, err = s.cacheRepository.GetUserBalance(tx, userID)
				if err != nil {
					return err
				}
			}
		}

		newBalance := balance + amount
		err = s.cacheRepository.IncreaseUserBalance(tx, userID, amount)
		if err != nil {
			return err
		}
		err = s.accountRepository.UpdateAccountBalance(ctx, nil, userID, newBalance)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
