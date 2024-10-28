package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *TournamentService) GetTournamentFilters(ctx context.Context, tournamentID string) (*dto.GetTournamentFiltersDTORes, error) {
	filtersDTO, err := s.tournamentQuerier.GetTournamentFilters(ctx, tournamentID)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "tournament", "error when getting tournamets")
	}

	return filtersDTO, nil
}
