package services

import (
	ports "github.com/DBrange/didis-comp-bk/domains/user/ports/drivens"
)

type UserService struct {
	userQueryer ports.ForQueryingUser
}

func NewUserService(userQueryer ports.ForQueryingUser) *UserService {
	return &UserService{
		userQueryer: userQueryer,
	}
}
