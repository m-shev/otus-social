package post

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/segmentio/kafka-go"
	"log"
	"strconv"
)

type CreatedPostMessage struct {
	PostId    int   `json:"postId"`
	AuthorId  int   `json:"authorId"`
	Consumers []int `json:"consumers"`
}

type Service struct {
	cache            *redis.Client
	postCreateReader *kafka.Reader
	logger           *log.Logger
	postCacheLimit   int
}

func NewPostService(cache *redis.Client, createPostReader *kafka.Reader, logger *log.Logger) *Service {
	return &Service{
		cache:            cache,
		postCreateReader: createPostReader,
		logger:           logger,
		postCacheLimit:   10,
	}
}

func (s *Service) ConsumeCreateMessage() {
	for {
		msg, err := s.postCreateReader.ReadMessage(context.Background())

		if err != nil {
			s.logger.Print(err.Error())
		}

		var postMessage CreatedPostMessage

		err = json.Unmarshal(msg.Value, &postMessage)

		if err != nil {
			s.logger.Printf("cant unmarshal message: %s", err.Error())
		}

		s.logger.Printf("message -> post created, id -> %d authorId -> %d, consumers %v",
			postMessage.PostId, postMessage.AuthorId, postMessage.Consumers)

		s.handleMessage(postMessage)
	}
}

func (s *Service) handleMessage(m CreatedPostMessage) {
	for _, v := range m.Consumers {
		key := strconv.Itoa(v)
		val, err := s.cache.Get(key).Result()

		if err == redis.Nil {
			postIds := []int{m.PostId}
			s.initNewPostIds(key, postIds)
			continue
		} else if err != nil {
			s.logCantGetPostIds(key, err)
			return
		}

		s.addToCache(key, val, m.PostId)
	}
}

func (s *Service) addToCache(key string, val string, id int) {
	var postIds []int
	err := json.Unmarshal([]byte(val), &postIds)

	if err != nil {
		s.logCantUnmarshalPostIds(key, err)
		return
	}

	postIds = append(postIds, id)

	if len(postIds) > s.postCacheLimit {
		postIds = postIds[1:]
	}

	b, err := json.Marshal(postIds)

	if err != nil {
		s.logCantMarshalPostIds(key, err)
		return
	}

	err = s.cache.Set(key, string(b), 0).Err()

	if err != nil {
		s.logCantMarshalPostIds(key, err)
	}
}

func (s *Service) initNewPostIds(key string, postIds []int) {
	bytes, err := json.Marshal(postIds)

	if err != nil {
		s.logCantMarshalPostIds(key, err)
		return
	}

	err = s.cache.Set(key, string(bytes), 0).Err()

	if err != nil {
		s.logCantSetPostIds(key, err)
	}
}
