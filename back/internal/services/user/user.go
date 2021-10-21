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

func (s *Service) GetProfileById(id int) (Profile, error) {
	u, err := s.repository.GetById(id)

	if err != nil {
		return Profile{}, err
	}

	return Profile{
		Id:        u.Id,
		Name:      u.Name,
		Surname:   u.Surname,
		Age:       u.Age,
		City:      u.City,
		Interests: u.Interests,
	}, nil
}

func hashUserPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func isHashEqual(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}

	return true
}