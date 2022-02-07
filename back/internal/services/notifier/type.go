package notifier

import "github.com/m-shev/otus-social/back/internal/services/post"

type PostQueue interface {
	SendPostCreated(m MessagePostCreate) error
}

type MessagePostCreate struct {
	AuthorId  int
	Post      post.Post
	Consumers []int
}

type MessageRecreateCache struct {
	Step string
	Post post.Post
}
