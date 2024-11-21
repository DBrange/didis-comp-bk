package services

import (
	"context"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *TournamentService) RemoveCompetitorFromTournament(ctx context.Context, tournamentID, competitorID string) error {
	tournamentRegistrationID, err := s.tournamentQuerier.GetTournamentRegistrationByCompetitorAndTournamentID(ctx, tournamentID, competitorID)
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when registering a competitor")
	}

	err = s.tournamentQuerier.WithTransaction(ctx, func(sessCtx mongo.SessionContext) error {
		if err := s.tournamentQuerier.DeleteTournamentRegistration(ctx, tournamentRegistrationID); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when registering a competitor")
		}

		if err := s.tournamentQuerier.DecrementTotalCompetitorsInTournament(ctx, tournamentID); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when decrement tournament total competitor")
		}
		return nil
	})

	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when remove competitor from tournament")
	}

	return nil
}
