package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *TournamentService) GetMatchByID(ctx context.Context, matchID string) (*dto.GetMatchDTORes, error) {
	categoryID, err := s.tournamentQuerier.GetMatchCategoryID(ctx, matchID)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "tournament", "error when getting tournament competitors")
	}

	matchDTO, err := s.tournamentQuerier.GetMatchByID(ctx, matchID, categoryID)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "tournament", "error when getting groups of the match")
	}

	return matchDTO, nil
}
