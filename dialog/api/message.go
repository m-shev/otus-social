package api

import (
	"github.com/gin-gonic/gin"
	"github.com/m-shev/otus-social/dialog/internal/services/message"
	"net/http"
	"strconv"
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

func (a *Api) GetMessage(c *gin.Context) {
	messageId := c.Param("messageId")
	dialogId, err := parseDialogId(c)

	if err != nil {
		return
	}

	m, err := a.messageService.GetMessage(dialogId, messageId)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, m)
}

func (a *Api) GetMessageList(c *gin.Context) {
	p, err := parseMessageListParams(c)

	if err != nil {
		return
	}

	dialogId, err := parseDialogId(c)

	if err != nil {
		return
	}

	l, err := a.messageService.GetMessageList(dialogId, p)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, l)
}

func parseMessageListParams(c *gin.Context) (message.ListParams, error) {

	take, err := strconv.Atoi(c.Query("take"))

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
		return message.ListParams{}, err
	}

	skip, err := strconv.Atoi(c.Query("skip"))

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
		return message.ListParams{}, err
	}

	return message.ListParams{Take: take, Skip: skip}, err
}

func parseDialogId(c *gin.Context) (int64, error) {
	dialogId, err := strconv.ParseInt(c.Param("dialogId"), 10, 64)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
	}

	return dialogId, err
}
