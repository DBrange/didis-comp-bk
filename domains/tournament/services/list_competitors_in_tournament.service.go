package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *TournamentService) ListCompetitorsInTournament(
	ctx context.Context,
	tournamentID,
	lastID string, limit int,
) (*dto.GetCompetitorsInTournamentDTORes, error) {
	categoryID, err := s.tournamentQuerier.GetCategoryIDOfTournament(ctx, tournamentID)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "tournament", "error when getting tournament competitors")
	}

	competitorsDTO, err := s.tournamentQuerier.GetCompetitorsInTournament(ctx, tournamentID, categoryID, lastID, limit, false)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "tournament", "error when getting tournament competitors")
	}

	TotalCompetitors, err := s.tournamentQuerier.GetTournamentTotalCompetitors(ctx, tournamentID)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "tournament", "error when getting tournament competitors")
	}

	competitors := &dto.GetCompetitorsInTournamentDTORes{
		Competitors: competitorsDTO,
		Total:       TotalCompetitors,
	}

	return competitors, nil
}
