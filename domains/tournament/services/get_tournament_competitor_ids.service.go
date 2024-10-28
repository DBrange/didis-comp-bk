package services

import (
	"context"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *TournamentService) GetTournamentCompetitorIDs(ctx context.Context, tournamentID string) ([]string, error) {
	competitorIDs, err := s.tournamentQuerier.GetTournamentCompetitorIDs(ctx, tournamentID)
	if err != nil {
		return []string{},customerrors.HandleErrMsg(err, "tournament", "error when getting competitorIDs")
	}

	return competitorIDs, nil
}
