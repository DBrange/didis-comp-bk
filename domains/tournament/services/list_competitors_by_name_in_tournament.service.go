package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *TournamentService) ListCompetitorsByNameInTournament(
	ctx context.Context,
	tournamentID string,
	name string,
	limit int,
) ([]*dto.GetCompetitorsInTournamentCompetitorDTORes, error) {
	categoryID, err := s.tournamentQuerier.GetCategoryIDOfTournament(ctx, tournamentID)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "tournament", "error when getting tournament competitors")
	}

	competitorsDTO, err := s.tournamentQuerier.GetCompetitorsByNameInTournament(ctx, tournamentID, categoryID, name, limit)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "tournament", "error when getting tournament competitors")
	}

	return competitorsDTO, nil
}
