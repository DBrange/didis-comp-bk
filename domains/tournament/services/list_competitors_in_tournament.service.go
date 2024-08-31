package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *TournamentService) ListCompetitorsInTournament(
	ctx context.Context,
	tournamentID, categoryID,
	lastID string, limit int,
) ([]*dto.GetCompetitorsInTournamentDTORes, error) {
	competitorsDTO, err := s.tournamentQuerier.GetCompetitorsInTournament(ctx, tournamentID, categoryID, lastID, limit)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "tournament", "error when getting tournament competitors")
	}

	return competitorsDTO, nil
}
