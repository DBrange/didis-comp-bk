package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
)

type ModifyPersonalInfo interface {
	ModifyPersonalInfo(ctx context.Context, userID string, userInfoDTO *dto.ModifyPersonalInfoDTOReq) error
}
