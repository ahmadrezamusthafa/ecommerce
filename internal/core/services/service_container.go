package services

import "github.com/ahmadrezamusthafa/ecommerce/internal/core/ports"

type ServiceContainer struct {
	UserService    ports.IUserService
	ProductService ports.IProductService
}

func NewServiceContainer(userService ports.IUserService,
	productService ports.IProductService) *ServiceContainer {
	return &ServiceContainer{
		UserService:    userService,
		ProductService: productService,
	}
}
