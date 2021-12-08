package main

import (
	"github.com/gin-gonic/gin"
	"github.com/m-shev/otus-social/dialog/internal/configuration"
	"log"
	"time"
)

func main() {
	config := configuration.GetConfig()
	time.Sleep(config.Server.StartDelay * time.Second)

	//logger := log.Default()
}

func router(conf config.Config, logger *log.Logger) *gin.Engine {
	//a := api.NewApi(conf.Db, logger, sessions.Default)

	handler := gin.New()

	handler.Use(gin.Logger())

	return handler
}
