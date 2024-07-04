package interfaces

import (
	user_dto "didis-comp-bk/internal/user/models/dto"
)

type CreateUser interface {
	CreateUser(user *user_dto.CreateUserDTO) error
}
