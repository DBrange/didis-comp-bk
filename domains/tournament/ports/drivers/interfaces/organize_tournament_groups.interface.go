package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

type OrganizeTournamentGroups interface {
	OrganizeTournamentGroups(ctx context.Context, tournamentID, roundID string, competitorDTOs []*dto.AddCompetitorsToTournamentGroupsDTOReq, sport models.SPORT,orderType, top int) error
}
