package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/location/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (driven *LocationService) CreateLocation(ctx context.Context, locationInfoDTO *dto.CreateLocationDTOReq) (string, error) {
	id, err := driven.locationQueryer.CreateLocation(ctx, locationInfoDTO)

	if err != nil {
		locationErrorHandlers := customerrors.CreateErrorHandlers("location")
		errMsgTemplate := "error inserting location"
		return "", customerrors.HandleError(err, locationErrorHandlers, errMsgTemplate)
	}

	return id, nil
}
