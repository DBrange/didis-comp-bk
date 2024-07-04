package interfaces

import (
	user_dto "github.com/DBrange/didis-comp-bk/internal/user/models/dto"
)

type GetUserByID interface {
	GetUserByID(id string) (*user_dto.GetUserByIDDTO, error)
}
