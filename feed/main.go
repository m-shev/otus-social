package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/m-shev/otus-social/feed/api"
	"github.com/m-shev/otus-social/feed/internal/configuration"
	"github.com/m-shev/otus-social/feed/internal/connector"
	"github.com/m-shev/otus-social/feed/internal/reader"
	"github.com/m-shev/otus-social/feed/internal/service/post"
	"log"
	"net/http"
)

func main() {
	logger := log.Default()
	config := configuration.GetConfig()

	consumerCache := connector.NewRedisConnector(config.Cache.Redis, config.Cache.ConsumerDb)
	postCache := connector.NewRedisConnector(config.Cache.Redis, config.Cache.PostDb)
	postReader := reader.NewPostReader(config.Broker, logger)
	postService := post.NewPostService(consumerCache, postCache, postReader, logger)
	postApi := api.NewApi(postService)
	handler := router(postApi)
	addr := fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)

	// start consuming create post messages
	go postService.ConsumeCreateMessage()

	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	logger.Printf("server start at %s", addr)

	if err := server.ListenAndServe(); err != nil {
		logger.Print("server start error: ", err)
	}
}

func router(api *api.Api) *gin.Engine {
	handler := gin.New()

	handler.Use(gin.Logger())

	handler.GET("/feed/:userId", api.GetFeed)
	return handler
}
