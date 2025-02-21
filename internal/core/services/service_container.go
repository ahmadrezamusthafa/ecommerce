package services

import "github.com/ahmadrezamusthafa/ecommerce/internal/core/ports"

type ServiceContainer struct {
	UserService ports.IUserService
}

func NewServiceContainer(userService ports.IUserService) *ServiceContainer {
	return &ServiceContainer{
		UserService: userService,
	}
}
