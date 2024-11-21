package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
)

type GetTournamentsByNameFromCategory interface {
GetTournamentsByNameFromCategory(ctx context.Context, categoryID string, sport models.SPORT, competitorType models.COMPETITOR_TYPE, tournamentName string) ([]*dto.GetTournamentsFromCategoryTournamentDTORes, error)
}
