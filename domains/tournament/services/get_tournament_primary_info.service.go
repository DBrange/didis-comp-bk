package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *TournamentService) GetTournamentPrimaryInfo(ctx context.Context, tournamentID string) (*dto.GetTournamentPrimaryInfoDTORes, error) {
	tournamentDTO, err := s.tournamentQuerier.GetTournamentPrimaryInfo(ctx, tournamentID)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "tournament", "error when getting tournamets")
	}

	return tournamentDTO, nil
}
