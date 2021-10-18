package api

import (
	"github.com/gin-gonic/gin"
	"github.com/m-shev/otus-social/internal/config"
	"github.com/m-shev/otus-social/internal/connector"
	user_storage "github.com/m-shev/otus-social/internal/repositories/user-storage"
	"github.com/m-shev/otus-social/internal/services/user"
	"log"
	"net/http"
)

type Api struct {
	userService *user.Service
}

func NewApi(conf config.Db, logger *log.Logger) *Api {
	dbConnect := connector.NewDbConnector(conf, logger)
	userRepository := user_storage.NewRepository(dbConnect)
	userService := user.NewService(userRepository)

	return &Api{
		userService: userService,
	}
}

func (a *Api) Registration(c *gin.Context) {
	var form user.CreateUserForm
	err := c.BindJSON(&form)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	u, err := a.userService.Create(form)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, u)
}
