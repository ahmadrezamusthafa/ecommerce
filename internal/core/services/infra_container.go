package services

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type InfraContainer struct {
	DB    *gorm.DB
	Cache *redis.Client
}

func NewInfraContainer(db *gorm.DB, cache *redis.Client) *InfraContainer {
	return &InfraContainer{
		DB:    db,
		Cache: cache,
	}
}
