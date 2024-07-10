//go:generate mockgen -destination=tests/mocks/mock_for_querying_user.go -package=mocks github.com/DBrange/didis-comp-bk/domains/user/ports/drivens ForQueryingUser

package ports

import (
	"context"

	dto_user "github.com/DBrange/didis-comp-bk/domains/user/models/dto"
)

type ForQueryingUser interface {
	CreateUser(ctx context.Context, userDTO *dto_user.CreateUserDTOReq) error
	GetUserByID(ctx context.Context, id string) (*dto_user.GetUserByIDDTO, error)
}
