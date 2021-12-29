package message

import "time"

type Message struct {
	MessageId string
	DialogId  int64
	AuthorId  int
	Content   string
	CreateAt  time.Time
}

type CreateMessageForm struct {
	DialogId int64
	AuthorId int
	Content  string
}

type ListParams struct {
	Take int
	Skip int
}
