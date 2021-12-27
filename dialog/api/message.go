package api

import (
	"github.com/gin-gonic/gin"
	"github.com/m-shev/otus-social/dialog/internal/services/message"
	"net/http"
)

func (a *Api) CreateMessage(c *gin.Context) {
	form := message.CreateMessageForm{}
	err := c.BindJSON(&form)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	m, err := a.messageService.CreateMessage(form)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, m)
}
