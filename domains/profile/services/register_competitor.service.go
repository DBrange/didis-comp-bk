package services

import (
	"context"
	"fmt"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/cmd/utils"
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *ProfileService) RegisterCompetitor(ctx context.Context, userIDs []string, sport models.SPORT, competitorType models.COMPETITOR_TYPE) error {
	err := s.profileQuerier.WithTransaction(ctx, func(sessCtx mongo.SessionContext) error {
		// Create type of competitor
		competiorTypeOID, err := s.CreateCompetitorType(sessCtx, competitorType)
		if err != nil {
			return nil
		}

		// Create competitor
		competitorID, err := s.profileQuerier.CreateCompetitor(sessCtx, sport, competitorType, competiorTypeOID)
		if err != nil {
			return err
		}

		switch len(userIDs) {
		case 1:
			err = s.registerCompetitorSingle(sessCtx, userIDs, competitorID)
		case 2:
			err = s.registerCompetitorDouble(sessCtx, userIDs, competitorID)
		default:
			err = fmt.Errorf("unsupported number of users: %d", len(userIDs))
		}

		return err
	})

	if err != nil {
		return customerrors.HandleErrMsg(err, "profile", "error when registering competitor")
	}

	return nil
}

func (s *ProfileService) registerCompetitorSingle(ctx context.Context, userIDs []string, competitorID string) error {
	userID := userIDs[0]
	// Get availability
	availabilitySliceDTO, err := s.profileQuerier.GetAvailabilityDailySlice(ctx, userID, "")
	if err != nil {
		return err
	}

	availabilitySliceOrder := utils.OrderAvailability(availabilitySliceDTO)

	// Create availability
	err = s.profileQuerier.CreateAvailabilityForCompetitor(ctx, competitorID, availabilitySliceOrder)
	if err != nil {
		return err
	}

	// Create competitor stats
	if err := s.profileQuerier.CreateCompetitorStats(ctx, competitorID); err != nil {
		return err
	}

	// Create competitor_user
	if err := s.profileQuerier.CreateCompetitorUser(ctx, userID, competitorID); err != nil {
		return err
	}

	return nil
}

func (s *ProfileService) registerCompetitorDouble(ctx context.Context, userIDs []string, competitorID string) error {
	// Availability users
	usersAvailabilitySliceDTO := make([][]*models.GetDailyAvailabilityByIDDTORes, len(userIDs))

	// Get availability users
	for i, userID := range userIDs {
		// Get availability
		availabilitySliceDTO, err := s.profileQuerier.GetAvailabilityDailySlice(ctx, userID, "")
		if err != nil {
			return err
		}

		usersAvailabilitySliceDTO[i] = availabilitySliceDTO
	}

	// Get availability
	availabilitySliceDTO := utils.IntermediateAvailability(usersAvailabilitySliceDTO)

	// Create availability
	err := s.profileQuerier.CreateAvailabilityForCompetitor(ctx, competitorID, availabilitySliceDTO)
	if err != nil {
		return err
	}

	// Create competitor stats
	if err := s.profileQuerier.CreateCompetitorStats(ctx, competitorID); err != nil {
		return err
	}

	for _, userID := range userIDs {
		if err := s.profileQuerier.VerifyUserExists(ctx , userID); err != nil{
			return err
		}
		// Create competitor_user
		if err := s.profileQuerier.CreateCompetitorUser(ctx, userID, competitorID); err != nil {
			return err
		}
	}

	return nil
}

func (r *ProfileService) CreateCompetitorType(ctx context.Context, competitorType models.COMPETITOR_TYPE) (string, error) {
	type createTypeCompetitor func(ctx context.Context) (string, error)

	createMap := map[models.COMPETITOR_TYPE]createTypeCompetitor{
		models.COMPETITOR_TYPE_SINGLE: func(ctx context.Context) (string, error) {
			singleDTO := &dto.CreateSingleDTOReq{}
			return r.profileQuerier.CreateSingle(ctx, singleDTO)
		},
		models.COMPETITOR_TYPE_DOUBLE: func(ctx context.Context) (string, error) {
			doubleDTO := &dto.CreateDoubleDTOReq{}
			return r.profileQuerier.CreateDouble(ctx, doubleDTO)
		},
		// models.COMPETITOR_TYPE_TEAM: func(ctx context.Context) (string, error) {
		// 	teamDTO := &dto.CreateTeamDTOReq{}
		// 	teamDTO.Admins = []string{userID}
		// 	return r.profileQuerier.CreateTeam(ctx, teamDTO)
		// },
	}

	create, ok := createMap[competitorType]
	if !ok {
		err := fmt.Errorf("error competitor type no exists: %w", customerrors.ErrNotFound)
		return "", customerrors.HandleErrMsg(err, "profile", "error when registering competitor")
	}

	return create(ctx)
}
