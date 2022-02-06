package api

import "github.com/m-shev/otus-social/feed/internal/service/post"

type FeedRequest struct {
	PostList []post.Content `json:"postList"`
	Total    int            `json:"total"`
}
