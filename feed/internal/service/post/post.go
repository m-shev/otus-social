package post

import (
	"database/sql"
	"github.com/go-redis/redis"
	"github.com/segmentio/kafka-go"
	"log"
)

type Service struct {
	consumerCache       *redis.Client
	postCache           *redis.Client
	postCreateReader    *kafka.Reader
	recreateCacheReader *kafka.Reader
	logger              *log.Logger
	db                  *sql.DB
	postCacheLimit      int
	readCreatePost      bool
	finishedReading     chan struct{}
}

func NewPostService(consumerCache *redis.Client, postCache *redis.Client, createPostReader *kafka.Reader,
	logger *log.Logger) *Service {
	return &Service{
		consumerCache:    consumerCache,
		postCache:        postCache,
		postCreateReader: createPostReader,
		logger:           logger,
		postCacheLimit:   50,
		readCreatePost:   true,
	}
}
