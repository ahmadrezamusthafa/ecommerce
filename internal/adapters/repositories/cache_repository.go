package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/ports"
	"github.com/ahmadrezamusthafa/ecommerce/internal/shared/constants"
	"github.com/go-redis/redis/v8"
	"strconv"
)

type cacheRepository struct {
	client *redis.Client
}

func NewCacheRepository(client *redis.Client) ports.ICacheRepository {
	return &cacheRepository{
		client: client,
	}
}

func (r *cacheRepository) GetUserBalance(tx *redis.Tx, userID int) (float64, error) {
	key := getUserBalanceKey(userID)
	val, err := tx.Get(context.Background(), key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0, nil
		}
		return 0, err
	}
	return strconv.ParseFloat(val, 64)
}

func (r *cacheRepository) IncreaseUserBalance(tx *redis.Tx, userID int, amount float64) (float64, error) {
	key := getUserBalanceKey(userID)
	result, err := tx.IncrByFloat(context.Background(), key, amount).Result()
	return result, err
}

func (r *cacheRepository) DecreaseUserBalance(tx *redis.Tx, userID int, amount float64) (float64, error) {
	key := getUserBalanceKey(userID)
	result, err := tx.IncrByFloat(context.Background(), key, -amount).Result()
	return result, err
}

func (r *cacheRepository) SetUserBalance(tx *redis.Tx, userID int, balance float64) (bool, error) {
	key := getUserBalanceKey(userID)
	res := tx.SetNX(context.Background(), key, balance, constants.DefaultTTLUserBalance)
	return res.Val(), res.Err()
}

func (r *cacheRepository) WatchUserBalance(userID int, fn func(tx *redis.Tx) error) error {
	return r.client.Watch(context.Background(), func(tx *redis.Tx) error {
		return fn(tx)
	}, getUserBalanceKey(userID))
}

func getUserBalanceKey(userID int) string {
	return fmt.Sprintf("user_balance:%d", userID)
}
