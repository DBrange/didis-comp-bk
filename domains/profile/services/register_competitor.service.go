package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *ProfileService) RegisterCompetitor(ctx context.Context, userID string, sport models.SPORT, competitorType models.COMPETITOR_TYPE) error {
	err := s.profileQueryer.WithTransaction(ctx, func(sessCtx mongo.SessionContext) error {
		// Convert to OID
		userOID, err := s.profileQueryer.ConvertToObjectID(userID)
		if err != nil {
			return err
		}

		// Create type of competitor
		competiorTypeOID, err := s.profileQueryer.CreateCompetitorType(ctx, competitorType)
		if err != nil {
			return nil
		}

		// Create competitor
		competitorID, err := s.profileQueryer.CreateCompetitor(ctx, sport, competitorType, competiorTypeOID)
		if err != nil {
			return err
		}

		// Convert to OID
		competitorOID, err := s.profileQueryer.ConvertToObjectID(competitorID)
		if err != nil {
			return err
		}

		// Create availability
		if err := s.profileQueryer.CreateAvailability(ctx, nil, &competitorID); err != nil {
			return err
		}

		// Create competitor stats
		if err := s.profileQueryer.CreateCompetitorStats(ctx, competitorOID); err != nil {
			return err
		}

		// Create competitor_user
		if err := s.profileQueryer.CreateCompetitorUser(ctx, userOID, competitorOID); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return customerrors.HandleErrMsg(err, "profile", "error when registering competitor")

	}

	return nil
}
