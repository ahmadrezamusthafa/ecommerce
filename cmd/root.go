package cmd

import (
	"context"
	"fmt"
	"github.com/ahmadrezamusthafa/ecommerce/config"
	"github.com/ahmadrezamusthafa/ecommerce/internal/adapters/repositories"
	"github.com/ahmadrezamusthafa/ecommerce/internal/adapters/rest"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/session"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/services"
	"github.com/ahmadrezamusthafa/ecommerce/pkg/cache"
	"github.com/ahmadrezamusthafa/ecommerce/pkg/database"
	"github.com/ahmadrezamusthafa/ecommerce/pkg/logger"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "ecommerce",
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
}

func start() {
	cfg := config.GetConfig()
	sessionCfg := session.GetSessionConfig()
	db, err := database.NewPostgresqlDatabase(cfg.Database)
	if err != nil {
		logger.Fatalf("Failed connect to database | %v", err)
	}
	logger.Info("Connected to database successfully")

	redisCache := cache.NewRedis(cfg.Cache)
	defer func() {
		_ = redisCache.Close()
	}()
	_, err = redisCache.Ping(context.Background()).Result()
	if err != nil {
		logger.Errorf("Failed connect to cache | %v", err)
	} else {
		logger.Info("Connected to cache successfully")
	}

	userRepository := repositories.NewUserRepository(db)
	productRepository := repositories.NewProductRepository(db)
	cartRepository := repositories.NewCartRepository(db)
	orderRepository := repositories.NewOrderRepository(db)

	userService := services.NewUserService(sessionCfg, userRepository)
	productService := services.NewProductService(sessionCfg, productRepository)
	cartService := services.NewCartService(sessionCfg, cartRepository)
	orderService := services.NewOrderService(sessionCfg, orderRepository, cartService)

	serviceContainer := services.NewServiceContainer(userService, productService, cartService, orderService)

	httpRouter := rest.InitRouter(cfg, sessionCfg, serviceContainer)
	httpRouter.Run()
}
