package user

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Create(form CreateUserForm) (User, error) {
	_, err := s.FindUser(form.Email)

	if err != nil && !errors.Is(err, ErrorUserNotFound) {
		return User{}, err
	}

	p, err := hashUserPassword(form.Password)

	if err != nil {
		return User{}, err
	}

	form.Password = p
	u, err := s.repository.Create(form)

	if errors.Is(err, ErrorUserAlreadyCreated) {
		return User{}, fmt.Errorf("user with email: %s has already been created", form.Email)
	}

	return u, err
}

func (s *Service) FindUser(email string) (User, error) {
	u, err := s.repository.FindUser(email)

	return u, err
}

func hashUserPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}