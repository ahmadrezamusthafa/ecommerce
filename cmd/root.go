package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ahmadrezamusthafa/ecommerce/config"
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

	b, _ := json.Marshal(db)
	fmt.Println(string(b))
}
