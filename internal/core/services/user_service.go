package services

import "github.com/ahmadrezamusthafa/ecommerce/internal/core/ports"

type userService struct {
	userRepository ports.IUserRepository
}

func NewUserService(ruleRepository ports.IUserRepository) ports.IUserService {
	return &userService{
		userRepository: ruleRepository,
	}
}
