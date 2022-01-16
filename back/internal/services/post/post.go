package post

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Create(f CreatePostForm) (Post, error) {
	return s.repository.Create(f)
}

func (s *Service) GetById(id int) (Post, error) {
	return s.repository.GetById(id)
}

func (s *Service) GetList(ids []int) ([]Post, error) {
	return s.repository.GetList(ids)
}
