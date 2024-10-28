package services

import (
	"context"

	tournament_registration_dto "github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *TournamentService) AddCompetitorInTournament(ctx context.Context, tournamentRegistrationDTO *tournament_registration_dto.CreateTournamentRegistrationDTOReq) error {
	available, err := s.tournamentQuerier.VerifyTournamentsCapacity(ctx, tournamentRegistrationDTO.TournamentID)
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when verify if tournament exits")
	}
	if !available {
		return customerrors.HandleErrMsg(err, "tournament", "tournament max capacity")
	}

	if err := s.tournamentQuerier.VerifyCompetitorExists(ctx, tournamentRegistrationDTO.CompetitorID); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when verify if competitor exits")
	}

	if err := s.tournamentQuerier.CreateTournamentRegistration(ctx, tournamentRegistrationDTO); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when registering a competitor")
	}

	if err := s.tournamentQuerier.IncrementTotalCompetitorsInTournament(ctx, tournamentRegistrationDTO.TournamentID); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when creating a tournament registration")
	}

	return nil
}
