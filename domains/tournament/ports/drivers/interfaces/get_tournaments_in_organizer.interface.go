package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

type GetTournamentsInOrganizer interface {
	GetTournamentsInOrganizer(
	ctx context.Context,
	organizerID string,
	sport models.SPORT,
	limit int,
	lastID string,
) (*dto.GetUserTournamentsDTORes, error)
}
