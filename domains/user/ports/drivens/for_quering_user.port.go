//go:generate mockgen -destination=tests/mocks/mock_for_querying_user.go -package=mocks github.com/DBrange/didis-comp-bk/domains/user/ports/drivens ForQueryingUser

package ports

import (
	"context"

	user_dto "github.com/DBrange/didis-comp-bk/domains/user/models/dto"
)

type ForQueryingUser interface {
	RegisterUser(ctx context.Context, userInfoDTO *user_dto.RegisterUserDTOReq) error
}
