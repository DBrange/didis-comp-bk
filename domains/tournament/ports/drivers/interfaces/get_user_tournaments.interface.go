package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

type GetUserTournaments interface {
	GetUserTournaments(
		ctx context.Context,
		userID string,
		sport models.SPORT,
		limit int,
		lastID string,
	) (*dto.GetUserTournamentsDTORes, error)
}
