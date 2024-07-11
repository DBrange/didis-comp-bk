package interfaces

import (
	"context"

	user_dto "github.com/DBrange/didis-comp-bk/domains/user/models/dto"
)

type CreateUser interface {
	CreateUser(ctx context.Context, user *user_dto.CreateUserDTOReq) error
}
