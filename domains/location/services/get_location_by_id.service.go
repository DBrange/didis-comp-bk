package services

import (
	"context"
	"errors"
	"fmt"

	location_dto "github.com/DBrange/didis-comp-bk/domains/location/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (d *LocationService) GetLocationByID(ctx context.Context, id string) (*location_dto.GetLocationByIDDTORes, error) {
	locationDTO, err := d.locationQuerier.GetLocationByID(ctx, id)
	if err != nil {
		return nil, getLocationByIDHandleError(err)
	}

	return locationDTO, nil
}

type getLocationByIDErrorHandler func(error) customerrors.AppError

var getLocationByIDErrorHandlers = map[error]getLocationByIDErrorHandler{
	customerrors.ErrInvalidID: func(err error) customerrors.AppError {
		return customerrors.AppError{
			Code: customerrors.ErrCodeDuplicateKey,
			Msg:  fmt.Sprintf("invalid location id format: %v", err),
		}
	},
	customerrors.ErrNotFound: func(err error) customerrors.AppError {
		return customerrors.AppError{
			Code: customerrors.ErrCodeSchemaViolation,
			Msg:  fmt.Sprintf("error when searching for location: %v", err),
		}
	},
}

func getLocationByIDHandleError(err error) error {
	for knownErr, handler := range getLocationByIDErrorHandlers {
		if errors.Is(err, knownErr) {
			return handler(err)
		}
	}
	return fmt.Errorf("error when searching for location: %w", err)
}
