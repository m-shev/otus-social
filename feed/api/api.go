package api

import "github.com/m-shev/otus-social/feed/internal/configuration"

type Api struct {
}

func NewApi(config configuration.Configuration) *Api {
	return &Api{}
}
