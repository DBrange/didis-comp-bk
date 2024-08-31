package services

import (
	"context"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *TournamentService) ModifyPots(ctx context.Context, tournamentID, potID, competitorID string, add bool) error {
	// Verify if pot is in tournament
	if err := s.tournamentQuerier.VerifyTournamentPot(ctx, tournamentID, potID); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error verifying pot in tournament")
	}

	// Verify if competitor is in tournament
	if err := s.tournamentQuerier.VerifyCompetitorExistsInTournament(ctx, tournamentID, competitorID); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error verifying competitor in tournament")
	}

	if add {
		// Add competitor in pot
		if err := s.tournamentQuerier.AddCompetitorInPot(ctx, potID, competitorID); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when add competitor in pot")
		}

		return nil
	}

	// Remove competitor of pot
	if err := s.tournamentQuerier.RemoveCompetitorOfPot(ctx, potID, competitorID); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when remove competitor of pot")
	}

	return nil
}
