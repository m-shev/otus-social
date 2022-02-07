package post

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis"
	"strconv"
)

func (s *Service) RecreateCacheMessage() {
	for {
		//msg, err := s.recreateCacheReader.ReadMessage(context.Background())
		<-s.finishedReading

		// recreateCache
	}
}

func (s *Service) ConsumeCreateMessage() {
	for s.readCreatePost {
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
			postMessage.Post.Id, postMessage.AuthorId, postMessage.Consumers)

		s.handleCreatePost(postMessage)
	}

	s.finishedReading <- struct{}{}
}

func (s *Service) handleCreatePost(m CreatedPostMessage) {

	contentCache, err := s.getContentCache(m.Post)

	if err != nil {
		s.logCantGetPost(m.Post.Id, err)
		return
	}

	for _, v := range m.Consumers {
		key := strconv.Itoa(v)
		val, err := s.consumerCache.Get(key).Result()

		if err == redis.Nil {
			postIds := []int{m.Post.Id}
			initErr := s.initNewPostIds(key, postIds)

			if initErr == nil {
				contentCache.ConsumerCount++
			}

			continue
		} else if err != nil {
			s.logCantGetPostIds(key, err)
			return
		}

		addErr := s.addPostToFeed(key, val, m.Post.Id)

		if addErr != nil {
			s.logCantSetPostIds(key, addErr)
		}

		contentCache.ConsumerCount++
	}

	s.setContentCache(contentCache)
}

func (s *Service) getContentCache(post Content) (ContentCache, error) {
	key := strconv.Itoa(post.Id)
	val, err := s.postCache.Get(key).Result()

	if err == redis.Nil {
		return ContentCache{
			Content:       post,
			ConsumerCount: 0,
		}, nil
	} else if err != nil {
		return ContentCache{}, err
	}

	var contentCache ContentCache

	err = json.Unmarshal([]byte(val), &contentCache)

	if err != nil {
		return ContentCache{}, err
	}

	return contentCache, err
}

func (s *Service) setContentCache(contentCache ContentCache) {
	key := strconv.Itoa(contentCache.Id)
	b, err := json.Marshal(contentCache)

	if err != nil {
		s.logCantMarshalPost(contentCache.Id, err)
	}

	err = s.postCache.Set(key, b, 0).Err()

	if err != nil {
		s.logCantSetPost(contentCache.Id, err)
	}
}

func (s *Service) checkPostConsumersCount(id int) {
	key := strconv.Itoa(id)
	b, err := s.postCache.Get(key).Bytes()

	if err != nil {
		s.logCantGetPost(id, err)
	}

	var contentCache ContentCache

	err = json.Unmarshal(b, &contentCache)

	if err != nil {
		s.logCantUnmarshalPost(id, err)
	}

	contentCache.ConsumerCount--

	if contentCache.ConsumerCount <= 0 {
		_, err = s.postCache.Del(key).Result()

		if err != nil {
			s.logCantDelPost(id, err)
		}
	} else {
		b, err = json.Marshal(contentCache)
		if err != nil {
			s.logCantMarshalPost(id, err)
		}
		err = s.postCache.Set(key, b, 0).Err()

		if err != nil {
			s.logCantSetPost(id, err)
		}
	}
}

func (s *Service) addPostToFeed(key string, val string, id int) error {
	var postIds []int
	err := json.Unmarshal([]byte(val), &postIds)

	if err != nil {
		s.logCantUnmarshalPostIds(key, err)
		return err
	}

	postIds = append(postIds, id)

	if len(postIds) > s.postCacheLimit {
		diff := len(postIds) - s.postCacheLimit
		toRemove := postIds[:diff]

		for _, v := range toRemove {
			s.checkPostConsumersCount(v)
		}

		postIds = postIds[diff:]
	}

	b, err := json.Marshal(postIds)

	if err != nil {
		s.logCantMarshalPostIds(key, err)
		return err
	}

	err = s.consumerCache.Set(key, string(b), 0).Err()

	if err != nil {
		s.logCantMarshalPostIds(key, err)
	}

	return nil
}

func (s *Service) initNewPostIds(key string, postIds []int) error {
	bytes, err := json.Marshal(postIds)

	if err != nil {
		s.logCantMarshalPostIds(key, err)
		return err
	}

	err = s.consumerCache.Set(key, string(bytes), 0).Err()

	if err != nil {
		s.logCantSetPostIds(key, err)
		return err
	}

	return nil
}
