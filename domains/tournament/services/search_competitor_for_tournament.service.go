package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *TournamentService) SearchCompetitorForTournament(ctx context.Context, userID string, name string, sport models.SPORT, competitorType models.COMPETITOR_TYPE) ([]*dto.GetCompetitorFollowedDTORes, error) {
	//VERIFICAR SI EXISTE EL USER

	competitors, err := s.tournamentQuerier.GetCompetitorsFollowed(ctx, userID, name, sport, competitorType)
	if err != nil {
		return nil, customerrors.HandleErrMsg(err, "tournament", "error when creating a tournament registration")
	}

	return competitors, nil
}
