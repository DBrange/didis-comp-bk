package services

import (
	"context"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *TournamentService) StartTournament(ctx context.Context, tournamentID, userID string, competitorsInMatchMap map[string][]string) error { 
	if err := s.tournamentQuerier.VerifyTournamentExists(ctx, tournamentID); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error match doesn't exists")
	}

	for matchID, competitorIDs := range competitorsInMatchMap {
		if err := s.tournamentQuerier.VerifyMatchExists(ctx, matchID); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error match doesn't exists")
		}
		if err := s.tournamentQuerier.VerifyMultipleCompetitorsExists(ctx, competitorIDs); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error competitor doesn't exists")
		}
	}

	err := s.tournamentQuerier.WithTransaction(ctx, func(sessCtx mongo.SessionContext) error {
		if err := s.tournamentQuerier.UpdateTournamentStartDate(ctx, tournamentID); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when creating chat")
		}

		if err := s.createMatchChats(ctx, competitorsInMatchMap, userID); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when creating chat")
		}

		return nil
	})
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when creating tournament location")
	}

	return nil
}

func (s *TournamentService) createMatchChats(ctx context.Context, competitorsInMatchMap map[string][]string, userID string) error {
	for matchID, competitorIDs := range competitorsInMatchMap {
		if len(competitorIDs) != 2 {
			continue
		}

		for _, competitorID := range competitorIDs {
			if competitorID == "" {
				continue
			}
		}

		areUsers, err := s.tournamentQuerier.VerifyCompetitorIDInCompetitorUser(ctx, competitorIDs)
		if err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when creating chat")
		}

		if areUsers {
			if err := s.tournamentQuerier.CreateMatchChat(ctx, matchID, competitorIDs, userID); err != nil {
				return customerrors.HandleErrMsg(err, "tournament", "error when creating chat")
			}
		}
	}

	return nil
}
