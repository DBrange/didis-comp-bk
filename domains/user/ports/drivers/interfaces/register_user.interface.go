package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/user/models/dto"
)

type RegisterUser interface {
	RegisterUser(ctx context.Context, userInfoDTO *dto.RegisterUserDTOReq) error
}
