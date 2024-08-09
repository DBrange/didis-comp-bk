package services

import (
	"context"
	"fmt"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *CategoryService) AddGuestUserInCategory(ctx context.Context, categoryID string, guestUsersDTO []*dto.CreateGuestUserDTOReq, sport models.SPORT, competitorType models.COMPETITOR_TYPE) error {
	var guestUserIDs []string

	if err := s.categoryQueryer.VerifyCategoryExists(ctx, categoryID); err != nil {
		return customerrors.HandleErrMsg(err, "category", "error verify category exists")
	}

	for _, guestUserDTO := range guestUsersDTO {
		// Create guest user
		guestUserID, err := s.categoryQueryer.CreateGuestUser(ctx, guestUserDTO)
		if err != nil {
			return customerrors.HandleErrMsg(err, "category", "error when creating a guest user")
		}

		guestUserIDs = append(guestUserIDs, guestUserID)
	}

	// Create type of competitor
	competiorTypeID, err := s.CreateCompetitorType(ctx, competitorType)
	if err != nil {
		return customerrors.HandleErrMsg(err, "category", "error when creating a competitor type")
	}

	// Create competitor
	competitorID, err := s.categoryQueryer.CreateCompetitor(ctx, sport, competitorType, competiorTypeID)
	if err != nil {
		return customerrors.HandleErrMsg(err, "category", "error when creating a competitor")
	}

	if err := s.categoryQueryer.CreateCompetitorStats(ctx, competitorID); err != nil {
		return err
	}

	for _, guestUserID := range guestUserIDs {
		// Use guest user ID and competitor ID for create guest_competitor
		guestCompetitorDTO := &dto.CreateGuestCompetitorDTOReq{
			GuestUserID:  guestUserID,
			CompetitorID: competitorID,
		}
		s.categoryQueryer.CreateGuestCompetitor(ctx, guestCompetitorDTO)
	}
	// Add competitor in category
	categoryRegistrationDTO := &dto.CreateCategoryRegistrationDTOReq{
		CompetitorID: competitorID,
		CategoryID:   categoryID,
	}

	// Add competitor in category
	if err := s.categoryQueryer.CreateCategoryRegistration(ctx, categoryRegistrationDTO); err != nil {
		return customerrors.HandleErrMsg(err, "category", "error when creating categoryRegistration")
	}

	if err := s.categoryQueryer.IncrementTotalParticipants(ctx, categoryID); err != nil {
		return customerrors.HandleErrMsg(err, "category", "error when increment total participants")
	}

	return nil
}

func (r *CategoryService) CreateCompetitorType(ctx context.Context, competitorType models.COMPETITOR_TYPE) (string, error) {
	type createTypeCompetitor func(ctx context.Context) (string, error)

	createMap := map[models.COMPETITOR_TYPE]createTypeCompetitor{
		models.COMPETITOR_TYPE_SINGLE: func(ctx context.Context) (string, error) {
			singleDTO := &dto.CreateSingleDTOReq{}
			return r.categoryQueryer.CreateSingle(ctx, singleDTO)
		},
		models.COMPETITOR_TYPE_DOUBLE: func(ctx context.Context) (string, error) {
			doubleDTO := &dto.CreateDoubleDTOReq{}
			return r.categoryQueryer.CreateDouble(ctx, doubleDTO)
		},
		models.COMPETITOR_TYPE_TEAM: func(ctx context.Context) (string, error) {
			teamDTO := &dto.CreateTeamDTOReq{}
			teamDTO.Admins = []string{}
			return r.categoryQueryer.CreateTeam(ctx, teamDTO)
		},
	}

	create, ok := createMap[competitorType]
	if !ok {
		err := fmt.Errorf("error competitor type no exists: %w", customerrors.ErrNotFound)
		return "", customerrors.HandleErrMsg(err, "profile", "error when registering competitor")
	}

	return create(ctx)
}
