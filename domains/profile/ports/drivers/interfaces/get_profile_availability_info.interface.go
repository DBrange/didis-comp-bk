package interfaces

import (
	"context"

	profile_dto "github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
)

type GetProfileDailyAvailabilityByID interface {
	GetProfileDailyAvailabilityByID(ctx context.Context, userID string, day string) (*profile_dto.GetDailyAvailabilityByIDDTORes, error)
}
