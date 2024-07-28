package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *TournamentService) AddGuestUserInTournament(ctx context.Context, tournamentID string, guestUserDTO *dto.CreateGuestUserDTOReq, sport models.SPORT, competitorType models.COMPETITOR_TYPE) error {
	// Create guest user
	guestUserID, err := s.tournamentQueryer.CreateGuestUser(ctx, guestUserDTO)
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when creating a guest user")
	}
	// Create type of competitor
	competiorTypeID, err := s.tournamentQueryer.CreateCompetitorType(ctx, competitorType)
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when creating a competitor type")
	}

	// Create competitor
	competitorID, err := s.tournamentQueryer.CreateCompetitor(ctx, sport, competitorType, competiorTypeID)
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when creating a competitor")
	}

	// Add competitor in tournament
	tournamentRegistrationDTO := &dto.CreateTournamentRegistrationDTOReq{
		TournamentID: tournamentID,
		CompetitorID: competitorID,
	}
	if err := s.tournamentQueryer.CreateTournamentRegistration(ctx, tournamentRegistrationDTO); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when creating a tournament registration")
	}

	// Use guest user ID and competitor ID for create guest_competitor
	guestCompetitorDTO := &dto.CreateGuestCompetitorDTOReq{
		GuestUserID:  guestUserID,
		CompetitorID: competitorID,
	}
	s.tournamentQueryer.CreateGuestCompetitor(ctx, guestCompetitorDTO)

	return nil
}
