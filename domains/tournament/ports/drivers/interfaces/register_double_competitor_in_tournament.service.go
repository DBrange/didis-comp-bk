package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
)

type RegisterDoubleCompetitorInTournament interface {
	RegisterDoubleCompetitorInTournament(ctx context.Context,tournamentID string, userIDs []string, sport models.SPORT, competitorType models.COMPETITOR_TYPE) error
}
