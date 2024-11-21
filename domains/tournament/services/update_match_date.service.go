package services

import (
	"context"
	"time"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *TournamentService) UpdateMatchDate(ctx context.Context, matchID string, date *time.Time) error {
	if err := s.tournamentQuerier.UpdateMatchDate(ctx, matchID, date); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when set date to match")
	}

	return nil
}
