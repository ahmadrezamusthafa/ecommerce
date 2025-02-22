package ports

import "github.com/go-redis/redis/v8"

type ICacheRepository interface {
	WatchUserBalance(userID string, fn func(tx *redis.Tx) error) error
}
