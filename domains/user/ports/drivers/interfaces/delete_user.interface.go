package interfaces

import (
	"context"

	user_dto "github.com/DBrange/didis-comp-bk/domains/user/models/dto"
)

type DeleteUser interface {
	DeleteUser(ctx context.Context, userID string) (*user_dto.UserRelationsToDeleteDTO, error)
}
