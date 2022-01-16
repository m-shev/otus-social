package api

import (
	"github.com/gin-gonic/gin"
	"github.com/m-shev/otus-social/back/internal/services/post"
	"net/http"
	"strconv"
)

func (a *Api) CreatePost(c *gin.Context) {
	var form post.CreatePostForm
	err := c.BindJSON(&form)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	p, err := a.postService.Create(form)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, p)
}

func (a *Api) GetById(c *gin.Context) {
	postId, err := strconv.ParseInt(c.Param("postId"), 10, 32)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	p, err := a.postService.GetById(int(postId))

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, p)
}

func (a *Api) GetPostList(c *gin.Context) {
	ids, _ := c.GetQueryArray("ids")
	idList, err := parsePostIds(ids)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	postList, err := a.postService.GetList(idList)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, postList)
}

func parsePostIds(ids []string) ([]int, error) {
	idList := make([]int, 0, len(ids))

	for _, v := range ids {
		postId, err := strconv.Atoi(v)

		if err != nil {
			return idList, err
		}

		idList = append(idList, postId)
	}

	return idList, nil
}
