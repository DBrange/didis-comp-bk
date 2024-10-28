package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

type ListCompetitorsByNameInTournament interface {
	ListCompetitorsByNameInTournament(
		ctx context.Context,
		tournamentID string,
		name string,
		limit int,
	) ([]*dto.GetCompetitorsInTournamentCompetitorDTORes, error)
}
