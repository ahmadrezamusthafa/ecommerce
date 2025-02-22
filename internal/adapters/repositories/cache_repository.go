package repositories

import (
	"context"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/ports"
	"github.com/go-redis/redis/v8"
)

type cacheRepository struct {
	client *redis.Client
}

func NewCacheRepository(client *redis.Client) ports.ICacheRepository {
	return &cacheRepository{
		client: client,
	}
}

func (r *cacheRepository) WatchUserBalance(userID string, fn func(tx *redis.Tx) error) error {
	return r.client.Watch(context.Background(), func(tx *redis.Tx) error {
		return fn(tx)
	}, getUserBalanceKey(userID))
}

func getUserBalanceKey(userID string) string {
	return "user_balance:" + userID
}
