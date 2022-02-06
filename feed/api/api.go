package api

import (
	"github.com/gin-gonic/gin"
	"github.com/m-shev/otus-social/feed/internal/service/post"
	"net/http"
	"strconv"
)

type Api struct {
	postService *post.Service
}

func NewApi(postService *post.Service) *Api {
	return &Api{
		postService: postService,
	}
}

func (a *Api) GetFeed(c *gin.Context) {
	params, err := parseFeedParams(c)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	postList, total, err := a.postService.GetFeed(params)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, FeedRequest{
		PostList: postList,
		Total:    total,
	})
}

func parseFeedParams(c *gin.Context) (post.FeedParams, error) {
	idStr := c.Param("userId")
	id, err := strconv.ParseInt(idStr, 10, 32)

	if err != nil {
		return post.FeedParams{}, err
	}

	takeStr := c.Query("take")
	take, err := strconv.ParseInt(takeStr, 10, 32)

	if err != nil {
		return post.FeedParams{}, err
	}

	skipStr := c.Query("skip")
	skip, err := strconv.ParseInt(skipStr, 10, 32)

	if err != nil {
		return post.FeedParams{}, err
	}

	return post.FeedParams{
		UserId: int(id),
		Take:   int(take),
		Skip:   int(skip),
	}, nil
}
