package services

import (
	"context"
	"errors"
	"fmt"

	location_dto "github.com/DBrange/didis-comp-bk/domains/location/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (d *LocationService) GetLocationByID(ctx context.Context, id string) (*location_dto.GetLocationByIDDTORes, error) {
	locationDTO, err := d.locationQueryer.GetLocationByID(ctx, id)
	if err != nil {
		if errors.Is(err, customerrors.ErrLocationNotFound) {
			appErr := customerrors.AppError{
				Code: customerrors.ErrCodeNotFound,
				Msg:  "error getting location: id not exists",
			}
			return nil, appErr
		}
		return nil, fmt.Errorf("error getting location: %w", err)
	}

	return locationDTO, nil
}