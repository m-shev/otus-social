package api

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/m-shev/otus-social/back/internal/services/user"
	"net/http"
)

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
func (a *Api) getUserIdFromSession(c *gin.Context) int {
	session := a.defaultSession(c)
	userId := session.Get("id")

	if userId != nil {
		return userId.(int)
	}

	return 0
}
