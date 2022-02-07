package api

import (
	"github.com/gin-gonic/gin"
	"github.com/m-shev/otus-social/back/internal/services/user"
	"net/http"
	"strconv"
)

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

func (a *Api) ListRegistration(c *gin.Context) {
	var forms []user.CreateUserForm
	err := c.BindJSON(&forms)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	for _, v := range forms {
		_, err := a.userService.Create(v)

		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
	}

	c.Status(http.StatusOK)
}

func (a *Api) MyProfile(c *gin.Context) {
	userId := a.getUserIdFromSession(c)

	if userId == 0 {
		c.Status(http.StatusUnauthorized)
		return
	}

	u, err := a.userService.GetProfileById(userId)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, u)
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

func (a *Api) AddFriend(c *gin.Context) {
	var addForm user.FriendForm
	err := c.BindJSON(&addForm)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	userId := a.getUserIdFromSession(c)

	if userId != addForm.UserId {
		c.String(http.StatusUnauthorized, "you must be logged in to add friends")
		c.Abort()
		return
	}

	err = a.userService.AddFriend(addForm)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	c.Status(http.StatusOK)
}

func (a *Api) RemoveFriend(c *gin.Context) {
	var removeFriend user.FriendForm
	err := c.BindJSON(&removeFriend)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	userId := a.getUserIdFromSession(c)

	if userId != removeFriend.UserId {
		c.String(http.StatusUnauthorized, "you must be logged in to add friends")
		c.Abort()
		return
	}

	err = a.userService.RemoveFriend(removeFriend)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	c.Status(http.StatusOK)
}

func (a *Api) FriendList(c *gin.Context) {
	userId, skip, take, err := convFriendListParam(c)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	friendList, total, err := a.userService.GetFriendList(userId, take, skip)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}
	resp := friendListResponse{
		FriendList: friendList,
		Total:      total,
	}

	c.JSON(http.StatusOK, resp)
}

func (a *Api) UserList(c *gin.Context) {
	name, surname, skip, take, err := convUserListParam(c)

	form := user.FindUsersForm{
		Name:    name,
		Surname: surname,
		Skip:    skip,
		Take:    take,
	}

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	userList, total, err := a.userService.FindUsers(form)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	res := userListResponse{
		UserList: userList,
		Total:    total,
	}

	c.JSON(http.StatusOK, res)
}

func convFriendListParam(c *gin.Context) (userId, skip, take int, err error) {
	userId, err = strconv.Atoi(c.Param("profileId"))

	if err != nil {
		return
	}

	skip, take, err = convPaginationParams(c)

	return
}

func convUserListParam(c *gin.Context) (name, surname string, skip, take int, err error) {
	skip, take, err = convPaginationParams(c)
	name = c.Query("name")
	surname = c.Query("surname")

	return
}

func convPaginationParams(c *gin.Context) (skip, take int, err error) {
	skip, err = strconv.Atoi(c.Query("skip"))

	if err != nil {
		return 0, 0, err
	}

	take, err = strconv.Atoi(c.Query("take"))

	return skip, take, err
}
