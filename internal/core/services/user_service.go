package services

import (
	"context"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/entity"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/ports"
)

type userService struct {
	userRepository ports.IUserRepository
}

func NewUserService(userRepository ports.IUserRepository) ports.IUserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	return s.userRepository.Create(ctx, user)
}

func (s *userService) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	return s.userRepository.Update(ctx, user)
}
