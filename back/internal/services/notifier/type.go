package notifier

type PostQueue interface {
	SendPostCreated(m MessagePostCreate) error
}

type MessagePostCreate struct {
	AuthorId  int
	PostId    int
	Consumers []int
}
