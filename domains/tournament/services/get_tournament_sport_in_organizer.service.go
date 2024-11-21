package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *TournamentService) GetTournamentSportsInOrganizer(ctx context.Context, organizerID string) ([]models.SPORT, error) {
	sports, err := s.tournamentQuerier.GetTournamentSportsInOrganizer(ctx, organizerID)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "tournament", "error when getting sports tournaments in organizer")
	}

	return sports, nil
}
