package api

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/m-shev/otus-social/internal/config"
	"github.com/m-shev/otus-social/internal/connector"
	user_storage "github.com/m-shev/otus-social/internal/repositories/user-storage"
	"github.com/m-shev/otus-social/internal/services/user"
	"log"
	"net/http"
	"strconv"
)

type DefaultSession = func (c *gin.Context)sessions.Session

type Api struct {
	userService    *user.Service
	defaultSession DefaultSession
}

func NewApi(conf config.Db, logger *log.Logger, defaultSession DefaultSession) *Api {
	dbConnect := connector.NewDbConnector(conf, logger)
	userRepository := user_storage.NewRepository(dbConnect)
	userService := user.NewService(userRepository)

	return &Api{
		userService:    userService,
		defaultSession: defaultSession,
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

func (a *Api) Auth(c *gin.Context) {
	var form user.AuthForm
	err := c.BindJSON(&form)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	u, err := a.userService.Auth(form)

	if err != nil {
		if errors.Is(err, user.ErrorUserUnauthorized) {
			c.String(http.StatusUnauthorized, err.Error())
		} else if errors.Is(err, user.ErrorUserNotFound) {
			c.String(http.StatusNotFound, err.Error())
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}

		return
	}

	session := a.defaultSession(c)
	session.Set("id", u.Id)
	session.Options(sessions.Options{
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	err = session.Save()

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, u)
}

func (a *Api) Logout(c *gin.Context) {
	session := a.defaultSession(c)
	session.Clear()
	err := session.Save()

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	c.Status(http.StatusNoContent)
}

func (a *Api) Profile(c *gin.Context) {
	profileId := c.Param("profileId")
	id, err := strconv.ParseInt(profileId, 10, 32)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	profile, err := a.userService.GetProfileById(int(id))

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, profile)
}