package services

import (
	"context"
	"fmt"

	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *TournamentService) UpdateQuantityGroupsInTournament(ctx context.Context, tournamentID string, position int, add bool) error {
	if add {
		if err := s.addGroupInTournament(ctx, tournamentID, position); err != nil {
			return err
		}
		return nil
	}

	if err := s.removeGroupToTournament(ctx, tournamentID, position); err != nil {
		return err
	}

	return nil
}

func (s *TournamentService) addGroupInTournament(ctx context.Context, tournamentID string, position int) error {
	if err := s.tournamentQueryer.VerifyNumberGroupsInTournament(ctx, tournamentID, position); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when set pot to tournament")
	}

	potID, err := s.tournamentQueryer.CreateTournamentGroup(ctx, tournamentID, position)
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when set pot to tournament")
	}

	if err := s.tournamentQueryer.AddGroupInTournament(ctx, tournamentID, potID); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when set pot to tournament")
	}

	return nil
}

func (s *TournamentService) removeGroupToTournament(ctx context.Context, tournamentID string, position int) error {
	// Get matches of the groups

	matchesToRemove, competitorIDs, err := s.tournamentQueryer.GetTournamentGroupMatchesByPosition(ctx, position, tournamentID)
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when adding competitors in groups")
	}
	fmt.Printf("partiodos a remover %v", matchesToRemove)

	if err := s.removeMatches(ctx, tournamentID, competitorIDs, matchesToRemove); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when eliminate group")
	}

	if len(matchesToRemove) > 0{
	if err := s.tournamentQueryer.DeleteMultipleMatches(ctx, matchesToRemove); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when creating competitorMatches")
	}

		if err := s.tournamentQueryer.DeleteMultipleCompetitorMatches(ctx, matchesToRemove); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when creating competitorMatches")
		}
	}

	if err := s.tournamentQueryer.RemoveGroupToTournament(ctx, tournamentID, position); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when set pot to tournament")
	}


	if err := s.tournamentQueryer.DeleteGroupByPosition(ctx, position, tournamentID); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when set pot to tournament")
	}


	// Get pots to update their position
	groupPositions, err := s.tournamentQueryer.GetTournamentGroupPositions(ctx, tournamentID)
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when getting pot positions in tournament")
	}

	newPotPostions := s.calculateNewPositions(groupPositions, position)


	for _, pp := range newPotPostions {
		if err := s.tournamentQueryer.UpdateGroupPositions(ctx, pp.ID, pp.Position); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when set pot to tournament")
		}
	}

	return nil
}
