package interfaces

import (
	"context"

	profile_dto "github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
)

type GetProfileAvailabilityInfoByID interface {
	GetProfileAvailabilityInfoByID(ctx context.Context, userID string, day string) (*profile_dto.GetProfileDailyAvailabilityInfoByIDDTORes, error)
}
