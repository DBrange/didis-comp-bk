package services

import (
	"context"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *ProfileService) CloseProfile(ctx context.Context, userID string) error {
	err := s.profileQuerier.WithTransaction(ctx, func(sessCtx mongo.SessionContext) error {
		userDeleted, err := s.profileQuerier.DeleteUser(ctx, userID)
		if err != nil {
			return err
		}

		locationID := userDeleted.LocationID

		availabilityID, err := s.profileQuerier.GetAvailabilityIDByUserID(ctx, userID)
		if err != nil {
			return err
		}

		err = s.profileQuerier.SetDeletedAt(ctx, s.profileQuerier.LocationColl(), locationID, "location")
		if err != nil {
			return err
		}

		err = s.profileQuerier.SetDeletedAt(ctx, s.profileQuerier.AvailabilityColl(), availabilityID, "availability")
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return customerrors.HandleErrMsg(err, "profile", "error updating deleted_at")
	}

	return nil
}
