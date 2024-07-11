package interfaces

import (
	"context"

	user_dto "github.com/DBrange/didis-comp-bk/domains/user/models/dto"
)

type GetUserByID interface {
	GetUserByID(ctx context.Context,id string) (*user_dto.GetUserByIDDTO, error)
}
