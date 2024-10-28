package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *TournamentService) GetTournamentAvailability(ctx context.Context, tournamentID string, day string) (*models.GetDailyAvailabilityByIDDTORes, string, error) {
	availabilityDTO, availabilityID, err := s.tournamentQuerier.GetDailyAvailabilityTournamentID(ctx, tournamentID, day)
	if err != nil {
		return nil, "", customerrors.HandleErrMsg(err, "tournament", "error getting profile info from category")
	}

	return availabilityDTO, availabilityID, nil
}
