package services

import (
	"context"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *TournamentService) ModifyRoundTotalPrize(ctx context.Context, roundID string, totalPrize float64) error {
	if err := s.tournamentQuerier.UpdateRoundTotalPrize(ctx, roundID, totalPrize); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when updating round total prize")
	}

	return nil
}
