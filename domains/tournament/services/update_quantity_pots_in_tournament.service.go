package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *TournamentService) UpdateQuantityPotsInTournament(ctx context.Context, tournamentID string, position int, add bool) error {
	if add {
		if err := s.addPotInTournament(ctx, tournamentID, position); err != nil {
			return err
		}
		return nil
	}

	if err := s.removePotToTournament(ctx, tournamentID, position); err != nil {
		return err
	}

	return nil

}

func (s *TournamentService) addPotInTournament(ctx context.Context, tournamentID string, position int) error {
	if err := s.tournamentQueryer.VerifyNumberPotsInTournament(ctx, tournamentID, position); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when set pot to tournament")
	}

	potID, err := s.tournamentQueryer.CreatePot(ctx, tournamentID, position)
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when set pot to tournament")
	}

	if err := s.tournamentQueryer.AddPotInTournament(ctx, tournamentID, potID); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when set pot to tournament")
	}

	return nil
}

func (s *TournamentService) removePotToTournament(ctx context.Context, tournamentID string, position int) error {
	if err := s.tournamentQueryer.RemovePotToTournament(ctx, tournamentID, position); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when set pot to tournament")
	}

	if err := s.tournamentQueryer.DeletePotByPosition(ctx, position, tournamentID); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when set pot to tournament")
	}

	// Get pots to update their position
	potPositions, err := s.tournamentQueryer.GetTournamentPotPositions(ctx, tournamentID)
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when getting pot positions in tournament")
	}

	newPotPostions := s.calculateNewPositions(potPositions, position)

	for _, pp := range newPotPostions {
		if err := s.tournamentQueryer.UpdatePotPositions(ctx, pp.ID, pp.Position); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when set pot to tournament")
		}
	}

	return nil
}

func (s *TournamentService) calculateNewPositions(positions []*dto.PotOrGroupPositionDTORes, positionToEliminate int) []*dto.PotOrGroupPositionDTORes {
	newPositions := make([]*dto.PotOrGroupPositionDTORes, len(positions))
	for i, position := range positions {

		if position.Position < positionToEliminate {
			newPositions[i] = &dto.PotOrGroupPositionDTORes{
				ID:       position.ID,
				Position: position.Position,
			}
		}
		if position.Position > positionToEliminate {
			newPositions[i] = &dto.PotOrGroupPositionDTORes{
				ID:       position.ID,
				Position: position.Position - 1,
			}

		}

	}

	return newPositions
}
