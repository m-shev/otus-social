package post

import "time"

type Repository interface {
	Create(form CreatePostForm) (Post, error)
	GetById(id int) (Post, error)
	GetList(ids []int) ([]Post, error)
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
