package services

import (
	"context"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *TournamentService) GetTournamentCompetitorIDsInMatches(ctx context.Context, tournamentID string) ([]string, error) {
	competitorIDs, err := s.tournamentQuerier.GetTournamentCompetitorIDsInMatches(ctx, tournamentID)
	if err != nil {
		return []string{},customerrors.HandleErrMsg(err, "tournament", "error when getting competitorIDs")
	}

	return competitorIDs, nil
}
