package services

import (
	"context"

	tournament_registration_dto "github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *TournamentService) AddCompetitorInTournament(ctx context.Context, tournamentRegistrationDTO *tournament_registration_dto.CreateTournamentRegistrationDTOReq) error {
	if err := s.tournamentQueryer.VerifyTournamentsExists(ctx, tournamentRegistrationDTO.TournamentID); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when verify if tournament exits")
	}

	if err := s.tournamentQueryer.VerifyCompetitorExists(ctx, tournamentRegistrationDTO.CompetitorID); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when verify if competitor exits")
	}

	if err := s.tournamentQueryer.CreateTournamentRegistration(ctx, tournamentRegistrationDTO); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when registering a competitor")
	}

	return nil
}
