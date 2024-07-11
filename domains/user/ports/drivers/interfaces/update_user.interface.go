package interfaces

import (
	"context"

	user_dto "github.com/DBrange/didis-comp-bk/domains/user/models/dto"
)

type UpdateUser interface {
	UpdateUser(ctx context.Context, userID string, user *user_dto.UpdateUserDTOReq) error
}
