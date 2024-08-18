package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *TournamentService) ModifyBracketMatch(ctx context.Context, matchID string, competitorsDTO []*dto.UpdateCompetitorMatchDTOReq) error {
	if err := s.tournamentQueryer.VerifyMatchExists(ctx, matchID); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error match doesn't exists")
	}

	for _, competitorDTO := range competitorsDTO {
		if err := s.tournamentQueryer.VerifyCompetitorExists(ctx, *competitorDTO.CompetitorID); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error competitor doesn't exists")
		}

		if err := s.tournamentQueryer.UpdateCompetitorMatch(ctx, matchID, competitorDTO); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when updating competitorMatch")
		}
	}

	return nil
}
