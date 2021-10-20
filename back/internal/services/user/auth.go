package user

func (s *Service) Auth(form AuthForm) (User, error)  {

	u, err := s.repository.FindUser(form.Login)

	if err != nil {
		return User{}, ErrorUserNotFound
	}

	hash, err := hashUserPassword(form.Password)

	if err != nil {
		return User{}, err
	}

	if !isHashEqual(hash, form.Password) {
		return User{}, ErrorUserUnauthorized
	}

	return u, nil
}