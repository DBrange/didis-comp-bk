package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
)

type GetProfileAvailabilityInCategory interface {
	GetProfileAvailabilityInCategory(ctx context.Context, competitorID, day string) (*models.GetDailyAvailabilityByIDDTORes,string, error)
}
