package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
)

type GetProfileAvailabilityInCategory interface {
	GetProfileAvailabilityInCategory(ctx context.Context, competitorID, day string) (*dto.GetDailyAvailabilityByIDDTORes, error)
}
