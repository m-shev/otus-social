package post

import (
	"strconv"
)

type Service struct {
	repository Repository
	queue      Queue
}

func NewService(repository Repository, queue Queue) *Service {
	return &Service{repository: repository, queue: queue}
}

func (s *Service) Create(f CreatePostForm) (Post, error) {
	post, err := s.repository.Create(f)

	if err != nil {
		return Post{}, err
	}

	err = s.queue.WriteJSON(strconv.Itoa(post.Id), CreatedPostMessage{
		PostId:   post.Id,
		AuthorId: post.AuthorId,
	})

	if err != nil {
		return post, ErrorSendMessageToQueue
	}

	return post, err
}

func (s *Service) GetById(id int) (Post, error) {
	return s.repository.GetById(id)
}

func (s *Service) GetList(params ListParams) ([]Post, error) {
	return s.repository.GetList(params)
}
