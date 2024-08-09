package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

type ListCompetitorsInTournament interface {
	ListCompetitorsInTournament(
		ctx context.Context,
		tournamentID,
		categoryID,
		lastID string,
		limit int,
	) ([]*dto.GetCompetitorsInTournamentDTORes, error)
}
