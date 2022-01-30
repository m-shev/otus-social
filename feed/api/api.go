package api

import (
	"github.com/m-shev/otus-social/feed/internal/configuration"
	"github.com/m-shev/otus-social/feed/internal/service/post"
)

type Api struct {
	postService *post.Service
}

func NewApi(config configuration.Configuration) *Api {
	return &Api{}
}
