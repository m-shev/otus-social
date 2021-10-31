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

type DefaultSession = func(c *gin.Context) sessions.Session

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

type MeetingService interface {
	CreateTask()
}

type CallService interface {
	CreateTask()
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

	friendList, err := a.userService.FindUsers(form)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, friendList)
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

func (a *Api) getUserIdFromSession(c *gin.Context) int {
	session := a.defaultSession(c)
	userId := session.Get("id")

	if userId != nil {
		return userId.(int)
	}

	return 0
}
