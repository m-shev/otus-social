package post

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
)

func (s *Service) GetFeed(p FeedParams) ([]Content, int, error) {
	key := strconv.Itoa(p.UserId)
	postIds, err := s.getPostIds(key)

	if err != nil {
		return []Content{}, 0, err
	}

	filtered := postPagination(postIds, p.Take, p.Skip)

	if len(filtered) > 0 {
		postList, err := s.getPostList(filtered)
		return postList, len(postIds), err
	}

	return []Content{}, 0, nil
}

func (s *Service) getPostIds(key string) ([]int, error) {
	b, err := s.consumerCache.Get(key).Bytes()

	if err != nil {
		return []int{}, err
	}

	var postIds []int

	err = json.Unmarshal(b, &postIds)
	return postIds, err
}

func (s *Service) getPostList(postIds []int) ([]Content, error) {
	postList := make([]Content, 0, len(postIds))
	keys := make([]string, 0, len(postIds))

	for _, v := range postIds {
		keys = append(keys, strconv.Itoa(v))
	}
	s.postCache.MGet()
	valList, err := s.postCache.MGet(keys...).Result()

	if err != nil {
		return postList, err
	}

	for _, v := range valList {
		var content ContentCache
		str := fmt.Sprintf("%v", v)
		err = json.Unmarshal([]byte(str), &content)

		if err == nil {
			postList = append(postList, content.Content)
		}

	}

	return postList, nil
}

func postPagination(postIds []int, take, skip int) []int {
	l := len(postIds)
	s := 0
	t := 0

	if take+skip >= l {
		s = 0
		t = l - skip
	} else if skip < l {

		s = l - skip - take
		t = take + s
	}

	res := postIds[s:t]
	_ = sort.Reverse(sort.IntSlice(res))

	return res
}
