package services

import (
	"context"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *TournamentService) AddCompetitorInTournamentGroup(ctx context.Context, groupID, tournamentID, competitorID string) error {
	// Verify competitor exists in tournament
	if err := s.tournamentQueryer.VerifyCompetitorExistsInTournament(ctx, tournamentID, competitorID); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when verifying competitors")
	}

	// Add competitor in tournament group
	if err := s.tournamentQueryer.AddCompetitorInGroup(ctx, groupID, competitorID); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when adding competitor in tournament group")
	}

	return nil

}
