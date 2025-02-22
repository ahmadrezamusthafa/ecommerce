package services

import "github.com/ahmadrezamusthafa/ecommerce/internal/core/ports"

type ServiceContainer struct {
	UserService    ports.IUserService
	ProductService ports.IProductService
	CartService    ports.ICartService
	OrderService   ports.IOrderService
	AccountService ports.IAccountService
}

func NewServiceContainer(userService ports.IUserService,
	productService ports.IProductService,
	cartService ports.ICartService,
	orderService ports.IOrderService,
	accountService ports.IAccountService) *ServiceContainer {
	return &ServiceContainer{
		UserService:    userService,
		ProductService: productService,
		CartService:    cartService,
		OrderService:   orderService,
		AccountService: accountService,
	}
}
