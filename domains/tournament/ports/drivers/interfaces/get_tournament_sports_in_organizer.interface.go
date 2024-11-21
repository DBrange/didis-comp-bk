package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
)

type GetTournamentSportsInOrganizer interface {
	GetTournamentSportsInOrganizer(ctx context.Context, organizerID string) ([]models.SPORT, error)
}
