package interfaces

import (
	"context"

	profile_dto "github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
)

type GetPersonalInfoByID interface {
	GetPersonalInfoByID(ctx context.Context, userID string) (*profile_dto.GetPersonalInfoByIDDTORes, error)
}
