package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

// import "context"

func (s *TournamentService) OrganizePots(ctx context.Context, tournamentID string, potDTOs []*dto.SetPotCompetitorDTOReq) error {
	potIDs := make([]string, len(potDTOs))
	var competitorIDs []string

	for i, potDTO := range potDTOs {
		potIDs[i] = potDTO.PotID
		competitorIDs = append(competitorIDs, potDTO.Competitors...)
	}

	// Verify if pots is in tournament
	if err := s.tournamentQueryer.VerifyMultipleTournamentPot(ctx, tournamentID, potIDs); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when add competitor in pot")
	}

	// Verify if competitors is in tournament
	if err := s.tournamentQueryer.VerifyMultipleCompetitorsExistsInTournament(ctx, tournamentID, competitorIDs); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when add competitor in pot")
	}

	// set competitors in pot
	if err := s.tournamentQueryer.SetCompetitorsInPots(ctx, tournamentID, potDTOs); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when add competitor in pot")
	}

	return nil
}
