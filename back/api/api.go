package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/m-shev/otus-social/back/internal/config"
	"github.com/m-shev/otus-social/back/internal/connector"
	"github.com/m-shev/otus-social/back/internal/repositories/broker"
	poststorage "github.com/m-shev/otus-social/back/internal/repositories/post-storage"
	userstorage "github.com/m-shev/otus-social/back/internal/repositories/user-storage"
	"github.com/m-shev/otus-social/back/internal/services/notifier"
	"github.com/m-shev/otus-social/back/internal/services/post"
	"github.com/m-shev/otus-social/back/internal/services/user"
	"log"
)

type DefaultSession = func(c *gin.Context) sessions.Session

type Api struct {
	userService    *user.Service
	postService    *post.Service
	notifier       *notifier.Service
	defaultSession DefaultSession
}

func NewApi(dbConf config.Db, brokerConf config.Broker, logger *log.Logger, defaultSession DefaultSession) *Api {
	dbConnect := connector.NewDbConnector(dbConf, logger)
	userRepository := userstorage.NewRepository(dbConnect)
	userService := user.NewService(userRepository)

	postRepository := poststorage.NewRepository(dbConnect)
	postQueue, err := broker.NewPostQueue(brokerConf, logger)

	if err != nil {
		panic(err)
	}

	postService := post.NewService(postRepository)

	notifierService := notifier.NewNotifierService(postQueue, userService, logger)

	return &Api{
		userService:    userService,
		postService:    postService,
		notifier:       notifierService,
		defaultSession: defaultSession,
	}
}
