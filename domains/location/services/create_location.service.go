package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/DBrange/didis-comp-bk/domains/location/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (driver *LocationService) CreateLocation(ctx context.Context, body *dto.CreateLocationDTOReq) (string, error) {
	id, err := driver.locationQueryer.CreateLocation(ctx, body)

	if err != nil {
		if errors.Is(err, customerrors.ErrLocationInsertionFailed) {
			appErr := customerrors.AppError{
				Code: customerrors.ErrCodeInsertionFailed,
				Msg:  fmt.Sprintf("error inserting location: %v", err),
			}
			return "", appErr
		}
		return "", fmt.Errorf("error inserting location: %w", err)
	}

	return id, nil
}
