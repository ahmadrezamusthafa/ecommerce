package services

import "github.com/ahmadrezamusthafa/ecommerce/internal/core/ports"

type ServiceContainer struct {
	UserService    ports.IUserService
	ProductService ports.IProductService
	CartService    ports.ICartService
}

func NewServiceContainer(userService ports.IUserService,
	productService ports.IProductService,
	cartService ports.ICartService) *ServiceContainer {
	return &ServiceContainer{
		UserService:    userService,
		ProductService: productService,
		CartService:    cartService,
	}
}
