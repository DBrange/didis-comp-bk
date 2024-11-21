package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
)

type GetProfileAvailabilityCompetitor interface {
	GetProfileAvailabilityCompetitor(ctx context.Context, competitorID, day string) (*models.GetDailyAvailabilityByIDDTORes, string, error)
}
