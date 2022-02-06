package post

import (
	"github.com/go-redis/redis"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

type CreatedPostMessage struct {
	Post      Content `json:"post"`
	AuthorId  int     `json:"authorId"`
	Consumers []int   `json:"consumers"`
}

type Content struct {
	Id        int       `json:"id"`
	AuthorId  int       `json:"authorId"`
	Content   string    `json:"content"`
	ImageLink string    `json:"imageLink"`
	CreateAt  time.Time `json:"createAt"`
	UpdateAt  time.Time `json:"updateAt"`
}

type ContentCache struct {
	Content
	ConsumerCount int `json:"consumerCount"`
}

type Service struct {
	consumerCache    *redis.Client
	postCache        *redis.Client
	postCreateReader *kafka.Reader
	logger           *log.Logger
	postCacheLimit   int
}

type FeedParams struct {
	UserId int
	Take   int
	Skip   int
}
