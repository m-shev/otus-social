package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/m-shev/otus-social/api"
	"github.com/m-shev/otus-social/internal/config"
	"github.com/m-shev/otus-social/internal/migration"
	"log"
	"net/http"
)

func main() {
	logger := log.Default()
	conf := config.GetConfig()
	m := migration.NewManager(conf.Db, logger)
	m.Up()

	router := makeRouter(conf, logger)
	server := &http.Server{
		Addr: "0.0.0.0:3005",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		logger.Print("server start error: ", err)
	}
}

func makeRouter(conf config.Config, logger *log.Logger) *gin.Engine {
	a := api.NewApi(conf.Db, logger)
	handler := gin.New()
	handler.Use(gin.Logger())
	handler.Use(cors.New(cors.Config{
		AllowAllOrigins:        false,
		AllowOrigins:          []string{"http://localhost:3000"},
	}))
	handler.POST("/user/registration", a.Registration)

	return handler
}