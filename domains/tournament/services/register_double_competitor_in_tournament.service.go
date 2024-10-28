package services

import (
	"context"
	"fmt"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/cmd/utils"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *TournamentService) RegisterDoubleCompetitorInTournament(ctx context.Context, tournamentID string, userIDs []string, sport models.SPORT, competitorType models.COMPETITOR_TYPE) error {
	err := s.tournamentQuerier.WithTransaction(ctx, func(sessCtx mongo.SessionContext) error {
		if err := s.tournamentQuerier.VerifyTournamentExists(ctx, tournamentID); err != nil {
			return err
		}

		if len(userIDs) != 2 {
			errLength := fmt.Errorf("unsupported number of users: %d", len(userIDs))
			return errLength
		}
		// Create type of competitor
		competiorTypeOID, err := s.CreateCompetitorType(sessCtx, competitorType)
		if err != nil {
			return nil
		}

		// Create competitor
		competitorID, err := s.tournamentQuerier.CreateCompetitor(sessCtx, sport, competitorType, competiorTypeOID)
		if err != nil {
			return err
		}

		if err = s.registerDoubleCompetitorInTournamentDouble(sessCtx, userIDs, competitorID); err != nil {
			return err
		}

		// Add competitor in tournament
		tournamentRegistrationDTO := &dto.CreateTournamentRegistrationDTOReq{
			TournamentID: tournamentID,
			CompetitorID: competitorID,
		}

		if err := s.tournamentQuerier.CreateTournamentRegistration(ctx, tournamentRegistrationDTO); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when creating a tournament registration")
		}

		if err := s.tournamentQuerier.IncrementTotalCompetitorsInTournament(ctx, tournamentID); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when creating a tournament registration")
		}

		return nil
	})

	if err != nil {
		return customerrors.HandleErrMsg(err, "profile", "error when registering competitor")
	}

	return nil
}

func (s *TournamentService) registerDoubleCompetitorInTournamentDouble(ctx context.Context, userIDs []string, competitorID string) error {
	// Availability users
	usersAvailabilitySliceDTO := make([][]*models.GetDailyAvailabilityByIDDTORes, len(userIDs))

	// Get availability users
	for i, userID := range userIDs {
		// Get availability
		availabilitySliceDTO, err := s.tournamentQuerier.GetAvailabilityDailySlice(ctx, userID, "")
		if err != nil {
			return err
		}

		usersAvailabilitySliceDTO[i] = availabilitySliceDTO
	}

	// Get availability
	availabilitySliceDTO := utils.IntermediateAvailability(usersAvailabilitySliceDTO)

	// Create availability
	err := s.tournamentQuerier.CreateAvailabilityForCompetitor(ctx, competitorID, availabilitySliceDTO)
	if err != nil {
		return err
	}

	// Create competitor stats
	if err := s.tournamentQuerier.CreateCompetitorStats(ctx, competitorID); err != nil {
		return err
	}

	for _, userOrGuestID := range userIDs {
		if err := s.tournamentQuerier.VerifyUserExists(ctx, userOrGuestID); err != nil {
			guestCompetitorDTO := &dto.CreateGuestCompetitorDTOReq{
				GuestUserID:  userOrGuestID,
				CompetitorID: competitorID,
			}
			s.tournamentQuerier.CreateGuestCompetitor(ctx, guestCompetitorDTO)
			break
		}
		// Create competitor_user
		if err := s.tournamentQuerier.CreateCompetitorUser(ctx, userOrGuestID, competitorID); err != nil {
			return err
		}
	}

	return nil
}
