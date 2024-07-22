package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
)

type RegisterUser interface {
	RegisterUser(ctx context.Context, profileInfoDTO *dto.RegisterUserDTOReq) error
}
