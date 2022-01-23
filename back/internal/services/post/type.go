package post

import "time"

type Repository interface {
	Create(form CreatePostForm) (Post, error)
	GetById(id int) (Post, error)
	GetList(params ListParams) ([]Post, error)
}

type Queue interface {
	WriteJSON(key string, v interface{}) error
}

type CreatedPostMessage struct {
	PostId   int `json:"postId"`
	AuthorId int `json:"authorId"`
}

type Post struct {
	Id        int       `json:"id"`
	AuthorId  int       `json:"authorId"`
	Content   string    `json:"content"`
	ImageLink string    `json:"imageLink"`
	CreateAt  time.Time `json:"createAt"`
	UpdateAt  time.Time `json:"updateAt"`
}

type CreatePostForm struct {
	AuthorId  int    `json:"authorId"`
	Content   string `json:"content"`
	ImageLink string `json:"imageLink"`
}

type ListParams struct {
	Ids      []int
	AuthorId int
	Take     int
	Skip     int
}
