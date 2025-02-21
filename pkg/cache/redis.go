package cache

import (
	"fmt"
	"github.com/ahmadrezamusthafa/ecommerce/config"
	"github.com/go-redis/redis/v8"
)

func NewRedis(cfg config.CacheConfiguration) *redis.Client {
	redisOptions := &redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		DialTimeout:  cfg.DialTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		PoolSize:     cfg.MaxActive,
		MinIdleConns: cfg.MaxIdle,
		MaxConnAge:   cfg.MaxLifetime,
	}

	if cfg.Password != "" {
		redisOptions.Password = cfg.Password
	}

	return redis.NewClient(redisOptions)
}
