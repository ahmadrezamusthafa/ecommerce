package services

import (
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/session"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/ports"
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
