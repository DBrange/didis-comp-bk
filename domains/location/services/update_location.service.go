package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/DBrange/didis-comp-bk/domains/location/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (driver *LocationService) UpdateLocation(ctx context.Context, locationID string, newLocationInfoDTO *dto.UpdateLocationDTOReq) error {
	err := driver.locationQueryer.UpdateLocation(ctx, locationID, newLocationInfoDTO)

	if err != nil {
		if errors.Is(err, customerrors.ErrInsertionFailed) {
			appErr := customerrors.AppError{
				Code: customerrors.ErrCodeInsertionFailed,
				Msg:  fmt.Sprintf("error inserting location: %v", err),
			}
			return appErr
		}
		return fmt.Errorf("error inserting location: %w", err)
	}

	return nil
}
