package user

func (s *Service) Auth(form AuthForm) (User, error)  {

	u, err := s.repository.FindUserByEmail(form.Login)

	if err != nil {
		return User{}, ErrorUserNotFound
	}

	if !isHashEqual(u.Password, form.Password) {
		return User{}, ErrorUserUnauthorized
	}

	return u, nil
}