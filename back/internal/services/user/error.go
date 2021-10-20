package user

import "errors"

var (
	ErrorUserNotFound = errors.New("user not found")
	ErrorUserAlreadyCreated = errors.New("user already created")
	ErrorUserUnauthorized = errors.New("user not authorized")
)
