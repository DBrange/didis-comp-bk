package services

import (
	"context"
	"errors"
	"fmt"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (driver *LocationService) DeleteLocation(ctx context.Context, locationID string) error {
	err := driver.locationQueryer.DeleteLocation(ctx, locationID)

	if err != nil {
		if errors.Is(err, customerrors.ErrLocationInsertionFailed) {
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
