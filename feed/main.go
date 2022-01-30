package main

import (
	"github.com/m-shev/otus-social/feed/internal/configuration"
	"github.com/m-shev/otus-social/feed/internal/connector"
	"github.com/m-shev/otus-social/feed/internal/reader"
	"github.com/m-shev/otus-social/feed/internal/service/post"
	"log"
)

func main() {
	logger := log.Default()
	config := configuration.GetConfig()

	cache := connector.NewRedisConnector(config.Redis)
	postReader := reader.NewPostReader(config.Broker, logger)
	postService := post.NewPostService(cache, postReader, logger)
	postService.ConsumeCreateMessage()
	//go produce(logger)
}
