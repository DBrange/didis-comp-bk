package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *TournamentService) GetRoundWithMatches(ctx context.Context, roundID, categoryID string) (*dto.GetRoundWithMatchesDTORes, error) {
	roundDTO, err := s.tournamentQuerier.GetRoundWithMatches(ctx, roundID, categoryID)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "tournament", "error when getting matches of the round")
	}

	return roundDTO, nil
}
