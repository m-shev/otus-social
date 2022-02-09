package connector

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/m-shev/otus-social/feed/internal/configuration"
	"time"
)

func NewRedisConnector(config configuration.Redis, db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		Password: "",
		DB:       db,
	})
}

func NewDbConnector(config configuration.Db) (*sql.DB, error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.User, config.Password, config.Host, config.Port, config.DbName)

	db, err := sql.Open("mysql", url)

	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * config.ConnMaxLifetime)
	db.SetMaxOpenConns(config.MaxOpenConnection)
	db.SetMaxIdleConns(config.MaxIdleConnection)

	return db, nil
}
