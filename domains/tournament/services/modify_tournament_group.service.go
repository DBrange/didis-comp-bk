package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *TournamentService) ModifyTournamentGroups(ctx context.Context, tournamentID, roundID string, competitorDTOs []*dto.AddCompetitorsToTournamentGroupsDTOReq, sport models.SPORT) error {
	// Verifications
	if err := s.verificationsModifyTournamentGroups(ctx, tournamentID, roundID, competitorDTOs); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when adding competitors in groups")
	}

	err := s.tournamentQuerier.WithTransaction(ctx, func(sessCtx mongo.SessionContext) error {
		for _, compecompetitorDTO := range competitorDTOs {
			// Get matches of the groups
			matchesToRemove, competitorIDs, err := s.tournamentQuerier.GetTournamentGroupMatches(ctx, compecompetitorDTO.GroupID)
			if err != nil {
				return customerrors.HandleErrMsg(err, "tournament", "error when adding competitors in groups")
			}

			if err := s.eliminateGroup(sessCtx, tournamentID, competitorIDs, matchesToRemove); err != nil {
				return customerrors.HandleErrMsg(err, "tournament", "error when eliminate group")
			}
		}

		// Add competitors on each group
		if err := s.tournamentQuerier.AddCompetitorsToTournamentGroups(sessCtx, tournamentID, competitorDTOs); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when adding competitors in groups")
		}

		// create matches alghoritm
		if err := s.createRoundRobinOneGroup(sessCtx, tournamentID, roundID, competitorDTOs, sport); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when adding competitors in groups")
		}

		return nil
	})
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when creating tournament location")
	}

	return nil
}

func (s *TournamentService) createRoundRobinOneGroup(ctx context.Context, tournamentID, roundID string, competitorDTOs []*dto.AddCompetitorsToTournamentGroupsDTOReq, sport models.SPORT) error {
	competitorsInMatchMap, err := s.createMatchesFromRoundRobin(ctx, tournamentID, roundID, competitorDTOs, sport)
	if err != nil {
		return err
	}

	courtAvailability, tournamentAvailabilities, timetablesNotAvailables, err := s.getCompleteAvailabilityInTournament(ctx, tournamentID)
	if err != nil {
		return err
	}

	if err := s.updateMatchesDates(ctx, competitorsInMatchMap, courtAvailability, tournamentAvailabilities, timetablesNotAvailables); err != nil {
		return err
	}

	return nil
}

func (s *TournamentService) verificationsModifyTournamentGroups(ctx context.Context, tournamentID, roundID string, competitorDTOs []*dto.AddCompetitorsToTournamentGroupsDTOReq) error {
	groupIDs := make([]string, len(competitorDTOs))
	var competitorIDs []string

	// Separating competitors of matches
	for i, competitorDTO := range competitorDTOs {
		groupIDs[i] = competitorDTO.GroupID
		competitorIDs = append(competitorIDs, competitorDTO.Competitors...)
	}

	// Verify if group is on tournament
	if err := s.tournamentQuerier.VerifyTournamentGroupInTournament(ctx, tournamentID, groupIDs); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when adding competitors in groups")
	}

	// Verify if round is on tournament
	if err := s.tournamentQuerier.VerifyRoundInTournament(ctx, roundID, tournamentID); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when adding competitors in groups")
	}

	// Verify if tournament contain all competitors
	if err := s.tournamentQuerier.VerifyMultipleCompetitorsExistsInTournament(ctx, tournamentID, competitorIDs); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when verifying competitors")
	}

	return nil
}

func (s *TournamentService) eliminateGroup(ctx context.Context, tournamentID string, competitorIDs, matchesToRemove []string) error {
	// sacar los matches de tournament
	if err := s.removeMatches(ctx, tournamentID, competitorIDs, matchesToRemove); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when adding competitors in groups")
	}

	if err := s.deleteMatches(ctx, matchesToRemove); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when adding competitors in groups")
	}

	return nil
}

func (s *TournamentService) removeMatches(ctx context.Context, tournamentID string, competitorIDs, matchesToRemove []string) error {
	// sacar los matches de tournament
	if err := s.tournamentQuerier.RemoveMultipleTournamentMatches(ctx, tournamentID, matchesToRemove); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when adding competitors in groups")
	}

	// sacar los matches de competitorStats y cambiar las
	if err := s.tournamentQuerier.RemoveMultipleCompetitorStatsMatches(ctx, competitorIDs, matchesToRemove); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when adding competitors in groups")
	}

	return nil
}

func (s *TournamentService) deleteMatches(ctx context.Context, matchesToRemove []string) error {
	if err := s.tournamentQuerier.DeleteMultipleMatches(ctx, matchesToRemove); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when creating competitorMatches")
	}

	if err := s.tournamentQuerier.DeleteMultipleCompetitorMatches(ctx, matchesToRemove); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when creating competitorMatches")
	}

	return nil
}
