package services

import (
	"context"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *ProfileService) CloseProfile(ctx context.Context, userID string) error {
	err := s.profileQueryer.WithTransaction(ctx, func(sessCtx mongo.SessionContext) error {
		userDeleted, err := s.profileQueryer.DeleteUser(ctx, userID)
		if err != nil {
			return err
		}

		locationID := userDeleted.LocationID

		availabilityID, err := s.profileQueryer.GetAvailabilityIDByUserID(ctx, userID)
		if err != nil {
			return err
		}

		err = s.profileQueryer.SetDeletedAt(ctx, s.profileQueryer.LocationColl(), locationID, "location")
		if err != nil {
			return err
		}

		err = s.profileQueryer.SetDeletedAt(ctx, s.profileQueryer.AvailabilityColl(), availabilityID, "availability")
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
