package services

import (
	"context"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *TournamentService) ModifyRoundPoints(ctx context.Context, roundID string, points int) error {
	if err := s.tournamentQuerier.UpdateRoundPoints(ctx, roundID, points); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when updating round points")
	}

	return nil
}
