package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

func (s *TournamentService) SearchCompetitorForTournament(ctx context.Context, userID string, name string, sport models.SPORT, competitorType models.COMPETITOR_TYPE) ([]*dto.GetCompetitorFollowedDTORes, error) {
	//VERIFICAR SI EXISTE EL USER
	
	return s.tournamentQuerier.GetCompetitorsFollowed(ctx, userID, name, sport, competitorType)
}
