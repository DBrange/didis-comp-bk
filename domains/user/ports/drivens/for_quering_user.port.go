//go:generate mockgen -destination=tests/mocks/mock_for_querying_user.go -package=mocks github.com/DBrange/didis-comp-bk/domains/user/ports/drivens ForQueryingUser

package ports

import (
	"context"

	user_dto "github.com/DBrange/didis-comp-bk/domains/user/models/dto"
)

type ForQueryingUser interface {
	CreateUser(ctx context.Context, userDTO *user_dto.CreateUserDTOReq) error
	GetUserByID(ctx context.Context, id string) (*user_dto.GetUserByIDDTO, error)
	UpdateUser(ctx context.Context,userID string, user *user_dto.UpdateUserDTOReq) error
}
