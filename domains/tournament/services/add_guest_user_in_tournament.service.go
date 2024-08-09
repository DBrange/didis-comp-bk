package services

import (
	"context"
	"fmt"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *TournamentService) AddGuestUserInTournament(ctx context.Context, tournamentID string, guestUsersDTO []*dto.CreateGuestUserDTOReq, sport models.SPORT, competitorType models.COMPETITOR_TYPE) error {
	var guestUserIDs []string

	if err := s.tournamentQueryer.VerifyTournamentsExists(ctx, tournamentID); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when creating a tournament registration")
	}

	for _, guestUserDTO := range guestUsersDTO {
		// Create guest user
		guestUserID, err := s.tournamentQueryer.CreateGuestUser(ctx, guestUserDTO)
		if err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when creating a guest user")
		}

		guestUserIDs = append(guestUserIDs, guestUserID)
	}

	// Create type of competitor
	competiorTypeID, err := s.CreateCompetitorType(ctx, competitorType)
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when creating a competitor type")
	}

	// Create competitor
	competitorID, err := s.tournamentQueryer.CreateCompetitor(ctx, sport, competitorType, competiorTypeID)
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when creating a competitor")
	}

	if err := s.tournamentQueryer.CreateCompetitorStats(ctx, competitorID); err != nil {
		return err
	}

	for _, guestUserID := range guestUserIDs {
		// Use guest user ID and competitor ID for create guest_competitor
		guestCompetitorDTO := &dto.CreateGuestCompetitorDTOReq{
			GuestUserID:  guestUserID,
			CompetitorID: competitorID,
		}
		s.tournamentQueryer.CreateGuestCompetitor(ctx, guestCompetitorDTO)
	}

	// Add competitor in tournament
	tournamentRegistrationDTO := &dto.CreateTournamentRegistrationDTOReq{
		TournamentID: tournamentID,
		CompetitorID: competitorID,
	}
	if err := s.tournamentQueryer.CreateTournamentRegistration(ctx, tournamentRegistrationDTO); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when creating a tournament registration")
	}

	return nil
}

func (r *TournamentService) CreateCompetitorType(ctx context.Context, competitorType models.COMPETITOR_TYPE) (string, error) {
	type createTypeCompetitor func(ctx context.Context) (string, error)

	createMap := map[models.COMPETITOR_TYPE]createTypeCompetitor{
		models.COMPETITOR_TYPE_SINGLE: func(ctx context.Context) (string, error) {
			singleDTO := &dto.CreateSingleDTOReq{}
			return r.tournamentQueryer.CreateSingle(ctx, singleDTO)
		},
		models.COMPETITOR_TYPE_DOUBLE: func(ctx context.Context) (string, error) {
			doubleDTO := &dto.CreateDoubleDTOReq{}
			return r.tournamentQueryer.CreateDouble(ctx, doubleDTO)
		},
		models.COMPETITOR_TYPE_TEAM: func(ctx context.Context) (string, error) {
			teamDTO := &dto.CreateTeamDTOReq{}
			teamDTO.Admins = []string{}
			return r.tournamentQueryer.CreateTeam(ctx, teamDTO)
		},
	}

	create, ok := createMap[competitorType]
	if !ok {
		err := fmt.Errorf("error competitor type no exists: %w", customerrors.ErrNotFound)
		return "", customerrors.HandleErrMsg(err, "profile", "error when registering competitor")
	}

	return create(ctx)
}
