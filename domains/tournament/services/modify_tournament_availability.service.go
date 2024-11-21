package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *TournamentService) ModifyTournamentAvailability(ctx context.Context, availabilityID string, availabilityInfoDTO *models.UpdateDailyAvailabilityDTOReq) error {
	if err := s.tournamentQuerier.UpdateAvailability(ctx, availabilityID, availabilityInfoDTO); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error updating tournament daily availability")
	}

	return nil
}