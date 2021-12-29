package api

import (
	"github.com/gin-gonic/gin"
	"github.com/m-shev/otus-social/dialog/internal/services/dialog"
	"net/http"
	"strconv"
)

func (a *Api) CreateDialog(c *gin.Context) {
	form := dialog.CreateDialogForm{}
	err := c.BindJSON(&form)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	d, err := a.dialogService.CreateDialog(form)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, d)
}

func (a *Api) GetDialogById(c *gin.Context) {
	dialogId := c.Param("dialogId")
	id, err := strconv.ParseInt(dialogId, 10, 64)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	d, err := a.dialogService.GetById(id)

	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, d)
}

func (a *Api) AddDialogMember(c *gin.Context) {
	form := dialog.AddMemberForm{}
	err := c.BindJSON(&form)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	form.Role = dialog.RoleMember

	err = a.dialogService.AddMember(form)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	c.Status(http.StatusOK)
}
