package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
)

type OrganizeTournamentGroups interface {
	OrganizeTournamentGroups(ctx context.Context, tournamentID, roundID string, sport models.SPORT, orderType, top, availableCourts, averageHours int) error
}
