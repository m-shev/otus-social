package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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
		Addr:    "0.0.0.0:3005",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		logger.Print("server start error: ", err)
	}
}

func makeRouter(conf config.Config, logger *log.Logger) *gin.Engine {
	a := api.NewApi(conf.Db, logger, sessions.Default)

	handler := gin.New()

	handler.Use(gin.Logger())

	handler.Use(cors.New(cors.Config{
		AllowAllOrigins:  false,
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowCredentials: true,
	}))

	//TODO secret to config
	store := cookie.NewStore([]byte("xxxx123"))

	//TODO session name to config
	handler.Use(sessions.Sessions("session", store))

	// router
	handler.POST("/user/registration", a.Registration)
	handler.POST("/user/auth", a.Auth)
	handler.POST("/user/friend", a.AddFriend)
	handler.POST("/user/list", a.UserList)
	handler.GET("/user/logout", a.Logout)
	handler.GET("/user/profile", a.MyProfile)
	handler.GET("/user/:profileId/profile", a.Profile)
	handler.GET("/user/:profileId/friends", a.FriendList)

	return handler
}
