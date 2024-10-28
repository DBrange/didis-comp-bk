package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *TournamentService) GetRoundGroups(ctx context.Context, roundID, categoryID string) (*dto.GetRoundGroupsDTORes, error) {
	roundDTO, err := s.tournamentQuerier.GetRoundGroups(ctx, roundID, categoryID)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "tournament", "error when getting groups of the round")
	}

	return roundDTO, nil
}
