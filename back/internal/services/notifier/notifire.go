package notifier

import (
	"github.com/m-shev/otus-social/back/internal/services/user"
	"log"
)

const (
	takeFriends = 50
)

type Service struct {
	userService *user.Service
	postQueue   PostQueue
	logger      *log.Logger
}

func NewNotifierService(postQueue PostQueue, userService *user.Service, logger *log.Logger) *Service {
	return &Service{
		userService: userService,
		postQueue:   postQueue,
		logger:      logger,
	}
}
