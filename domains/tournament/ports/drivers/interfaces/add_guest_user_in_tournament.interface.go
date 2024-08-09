package interfaces

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

type AddGuestUserInTournament interface {
	AddGuestUserInTournament(ctx context.Context, tournamentID string, guestUsersDTO []*dto.CreateGuestUserDTOReq, sport models.SPORT, competitorType models.COMPETITOR_TYPE) error
}
