package rest

import (
	"context"
	"errors"
	"fmt"
	"github.com/ahmadrezamusthafa/ecommerce/config"
	"github.com/ahmadrezamusthafa/ecommerce/internal/adapters/rest/handlers"
	"github.com/ahmadrezamusthafa/ecommerce/internal/adapters/rest/middlewares"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/session"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/services"
	"github.com/ahmadrezamusthafa/ecommerce/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Router struct {
	*gin.Engine
	cfg              *config.Configuration
	serviceContainer *services.ServiceContainer
}

func InitRouter(cfg *config.Configuration,
	sessionCfg *session.Config,
	serviceContainer *services.ServiceContainer) *Router {
	router := gin.Default()

	userHandler := handlers.NewUserHandler(serviceContainer)
	productHandler := handlers.NewProductHandler(serviceContainer)
	cartHandler := handlers.NewCartHandler(serviceContainer)
	orderHandler := handlers.NewOrderHandler(serviceContainer)
	accountHandler := handlers.NewAccountHandler(serviceContainer)

	apiV1 := router.Group("/api/v1")

	apiV1.POST("/user/register", userHandler.Register)
	apiV1.POST("/user/login", userHandler.Login)
	apiV1.PUT("/user/update", middlewares.AuthMiddleware(sessionCfg), userHandler.Update)

	apiV1.GET("/products", productHandler.GetAllProducts)
	apiV1.GET("/products/:id", productHandler.GetProductByID)
	apiV1.GET("/products/search", productHandler.SearchProducts)

	apiV1.GET("/cart", middlewares.AuthMiddleware(sessionCfg), cartHandler.GetCart)
	apiV1.POST("/cart/items", middlewares.AuthMiddleware(sessionCfg), cartHandler.AddItemToCart)
	apiV1.DELETE("/cart/items/:id", middlewares.AuthMiddleware(sessionCfg), cartHandler.RemoveItemFromCart)

	apiV1.POST("/orders", middlewares.AuthMiddleware(sessionCfg), orderHandler.CreateOrder)
	apiV1.GET("/orders", middlewares.AuthMiddleware(sessionCfg), orderHandler.GetOrder)
	apiV1.GET("/orders/top-customers", middlewares.AuthMiddleware(sessionCfg), orderHandler.GetTopCustomers)

	apiV1.POST("/accounts/deposit", middlewares.AuthMiddleware(sessionCfg), accountHandler.Deposit)
	apiV1.POST("/accounts/withdraw", middlewares.AuthMiddleware(sessionCfg), accountHandler.Withdraw)

	return &Router{
		Engine:           router,
		cfg:              cfg,
		serviceContainer: serviceContainer,
	}
}

func (router *Router) Run() {
	port := router.cfg.App.Port
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router.Handler(),
	}

	logger.Infof("Starting http server on port %d", port)

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatalf("Failed to serve server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("Server shutdown: %v", err)
	}
	select {
	case <-ctx.Done():
	}
	logger.Info("Server exiting")
}
