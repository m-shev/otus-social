package connector

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/m-shev/otus-social/feed/internal/configuration"
)

func NewRedisConnector(config configuration.Redis) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		Password: "",
		DB:       config.DB,
	})
}
