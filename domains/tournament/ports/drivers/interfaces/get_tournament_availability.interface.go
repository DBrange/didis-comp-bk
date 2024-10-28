package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
)

type GetTournamentAvailability interface {
	GetTournamentAvailability(ctx context.Context, tournamentID string, day string) (*models.GetDailyAvailabilityByIDDTORes, string,error)
}
