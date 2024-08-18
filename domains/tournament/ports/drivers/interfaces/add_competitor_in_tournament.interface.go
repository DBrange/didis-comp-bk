package interfaces

import (
	"context"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

type AddCompetitorInTournament interface {
	AddCompetitorInTournament(ctx context.Context, tournamentResgistrationDTO *dto.CreateTournamentRegistrationDTOReq) error
}
