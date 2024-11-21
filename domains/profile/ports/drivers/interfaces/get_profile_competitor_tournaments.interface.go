package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
)

type GetProfileCompetitorTournaments interface {
	GetProfileCompetitorTournaments(
		ctx context.Context,
		competitorID, categoryID string,
		sport models.SPORT,
		limit int,
		lastID string,
	) (*dto.GetProfileUserTournamentsDTORes, error)
}
