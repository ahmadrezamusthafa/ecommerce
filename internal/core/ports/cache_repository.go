package ports

import "github.com/go-redis/redis/v8"

type ICacheRepository interface {
	SetUserBalance(tx *redis.Tx, userID int, balance float64) (bool, error)
	IncreaseUserBalance(tx *redis.Tx, userID int, amount float64) (float64, error)
	DecreaseUserBalance(tx *redis.Tx, userID int, amount float64) (float64, error)
	GetUserBalance(tx *redis.Tx, userID int) (float64, error)
	WatchUserBalance(userID int, fn func(tx *redis.Tx) error) error
}
