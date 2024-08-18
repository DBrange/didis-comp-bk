package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *TournamentService) OrganizeTournamentGroups(ctx context.Context, tournamentID, roundID string, competitorDTOs []*dto.AddCompetitorsToTournamentGroupsDTOReq, sport models.SPORT) error {
	// Verifications
	if err := s.verificationsOrganizeTournamentGroups(ctx, tournamentID, roundID, competitorDTOs); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when adding competitors in groups")
	}

	err := s.tournamentQueryer.WithTransaction(ctx, func(sessCtx mongo.SessionContext) error {
		// Add competitors on each group
		if err := s.tournamentQueryer.AddCompetitorsToTournamentGroups(sessCtx, tournamentID, competitorDTOs); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when adding competitors in groups")
		}

		// create matches alghoritm
		if err := s.createRoundRobin(sessCtx, tournamentID, roundID, competitorDTOs, sport); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when adding competitors in groups")
		}

		return nil
	})
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when creating tournament location")
	}

	return nil
}

func (s *TournamentService) createRoundRobin(ctx context.Context, tournamentID, roundID string, competitorDTOs []*dto.AddCompetitorsToTournamentGroupsDTOReq, sport models.SPORT) error {
	for _, competitorDTO := range competitorDTOs {
		competitors := competitorDTO.Competitors
		n := len(competitors)
		matchesNum := s.CalculateTotalMatchesInGroup(competitors) // Número total de partidos en un sistema "todos contra todos"

		competitorMatchDTO := make([]*dto.UpdateCompetitorMatchDTOReq, matchesNum*2) // *2 porque cada partido tiene 2 competidores

		// Create matches
		matchIDs, err := s.createMatchesInGroup(ctx, matchesNum, roundID, tournamentID, sport)
		if err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when creating matches")
		}

		// Add matches in tournament
		if err := s.addMatchesInTournament(ctx, tournamentID, competitorDTO.GroupID, matchIDs); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when adding competitors in groups")
		}

		matchIndex := 0
		for i := 0; i < n-1; i++ {
			for j := i + 1; j < n; j++ {
				if matchIndex >= len(matchIDs) {
					break
				}
				matchID := matchIDs[matchIndex]

				// Competitor objects
				competitorMatchDTO[matchIndex*2] = &dto.UpdateCompetitorMatchDTOReq{
					MatchID:      &matchID,
					CompetitorID: &competitors[i],
					Position:     utils.IntPtr(1),
				}
				competitorMatchDTO[matchIndex*2+1] = &dto.UpdateCompetitorMatchDTOReq{
					MatchID:      &matchID,
					CompetitorID: &competitors[j],
					Position:     utils.IntPtr(2),
				}

				// Add competitors to stats
				if err := s.tournamentQueryer.AddMatchInCompetitorStats(ctx, competitors[i], matchID); err != nil {
					return customerrors.HandleErrMsg(err, "tournament", "error when adding match competitor stats")
				}
				if err := s.tournamentQueryer.AddMatchInCompetitorStats(ctx, competitors[j], matchID); err != nil {
					return customerrors.HandleErrMsg(err, "tournament", "error when adding match competitor stats")
				}

				matchIndex++
			}
		}

		// Update competitorMatches with the new info
		if err := s.tournamentQueryer.UpdateMultipleCompetitorMatches(ctx, competitorMatchDTO); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when updating competitorMatches")
		}
	}

	return nil
}

func (s *TournamentService) createMatchesInGroup(ctx context.Context, matchesNumber int, roundID, tournamentID string, sport models.SPORT) ([]string, error) {
	matchIDs := make([]string, matchesNumber)
	for i := 0; i < int(matchesNumber); i++ {
		matchDTO := &dto.CreateMatchDTOReq{
			Sport:        sport,
			TournamentID: tournamentID,
			RoundID:      roundID,
			Position:     0,
		}

		// Create match
		matchID, err := s.tournamentQueryer.CreateMatch(ctx, matchDTO)
		if err != nil {
			return []string{}, customerrors.HandleErrMsg(err, "tournament", "error when creating matches")
		}
		matchIDs[i] = matchID

		// Create competitorMatch
		if err := s.CreateCompetitorMatches(ctx, matchID); err != nil {
			return []string{}, customerrors.HandleErrMsg(err, "tournament", "error when creating competitorMatches")
		}
	}

	return matchIDs, nil
}

// CalculateTotalMatches calcula la cantidad de partidos que deben realizarse en un torneo de "todos contra todos".
func (s *TournamentService) CalculateTotalMatchesInGroup(competitorIDs []string) int {
	n := len(competitorIDs)
	// Combinaciones de n competidores tomados de 2 en 2: n * (n - 1) / 2
	totalMatches := (n * (n - 1)) / 2
	return totalMatches
}

func (s *TournamentService) verificationsOrganizeTournamentGroups(ctx context.Context, tournamentID, roundID string, competitorDTOs []*dto.AddCompetitorsToTournamentGroupsDTOReq) error {
	groupIDs := make([]string, len(competitorDTOs))
	var competitorIDs []string

	// Separating competitors of matches
	for i, competitorDTO := range competitorDTOs {
		groupIDs[i] = competitorDTO.GroupID
		competitorIDs = append(competitorIDs, competitorDTO.Competitors...)
	}

	// Verify if group is on tournament
	if err := s.tournamentQueryer.VerifyMultipleGroupsInTournament(ctx, tournamentID, groupIDs); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when adding competitors in groups")
	}

	// Verify if round is on tournament
	if err := s.tournamentQueryer.VerifyRoundInTournament(ctx, roundID, tournamentID); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when adding competitors in groups")
	}

	// Verify if tournament contain all competitors
	if err := s.tournamentQueryer.VerifyMultipleCompetitorsExistsInTournament(ctx, tournamentID, competitorIDs); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when verifying competitors")
	}

	return nil
}

func (s *TournamentService) addMatchesInTournament(ctx context.Context, tournamentID, groupID string, matchIDs []string) error {
	// Add matches in tournament
	if err := s.tournamentQueryer.AddMultipleMatchesInTournament(ctx, tournamentID, matchIDs); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when adding matches in tournament")
	}

	// Add matches in group
	if err := s.tournamentQueryer.AddMultipleMatchesInTournamentGroup(ctx, groupID, tournamentID, matchIDs); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when adding matches in group")
	}

	return nil
}
