package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
)

type Login interface {
	Login(ctx context.Context, loginDTO *dto.LoginDTOReq) (*dto.GetUserForLoginDTO, string, string, error)
}
