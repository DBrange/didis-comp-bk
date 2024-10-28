package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *ProfileService) GetOrganizerData(ctx context.Context, userID string) (*dto.GetOrganizerDataDTORes, error) {
	organizerDTO, err := s.profileQuerier.GetOrganizerData(ctx, userID)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "profile", "error getting organizer data")
	}

	return organizerDTO, nil
}
