package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/m-shev/otus-social/back/api"
	"github.com/m-shev/otus-social/back/internal/config"
	"github.com/m-shev/otus-social/back/internal/migration"
	"log"
	"net/http"
	"time"
)

func main() {
	conf := config.GetConfig()

	time.Sleep(conf.Server.StartDelay * time.Second)
	logger := log.Default()

	m := migration.NewManager(conf.Db, logger)
	m.Up()
	router := makeRouter(conf, logger)
	addr := fmt.Sprintf("%s:%s", conf.Server.Host, conf.Server.Port)

	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	logger.Printf("server start at %s", addr)

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
		AllowOrigins:     conf.AllowOrigins,
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "DELETE"},
	}))

	store := cookie.NewStore([]byte("IRdjDm"))

	handler.Use(sessions.Sessions("session", store))

	// USER
	handler.POST("/user/registration", a.Registration)
	handler.POST("user/registrations", a.ListRegistration)
	handler.POST("/user/auth", a.Auth)
	handler.POST("/user/friend", a.AddFriend)
	handler.DELETE("/user/friend", a.RemoveFriend)
	handler.GET("/user/list", a.UserList)
	handler.GET("/user/logout", a.Logout)
	handler.GET("/user/profile", a.MyProfile)
	handler.GET("/user/:profileId/profile", a.Profile)
	handler.GET("/user/:profileId/friends", a.FriendList)

	// POST
	handler.POST("/post", a.CreatePost)
	handler.GET("/post/:postId", a.GetById)
	handler.GET("/post", a.GetPostList)
	return handler
}
