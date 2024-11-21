package services

import (
	"context"
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *TournamentService) ModifyBracketMatch(ctx context.Context, tournamentID string, competitorDTOs []*dto.UpdateCompetitorMatchDTOReq) error {
	err := s.tournamentQuerier.WithTransaction(ctx, func(sessCtx mongo.SessionContext) error {
		for _, competitorDTO := range competitorDTOs {
			if err := s.tournamentQuerier.VerifyMatchExists(ctx, competitorDTO.MatchID); err != nil {
				return customerrors.HandleErrMsg(err, "tournament", "error match doesn't exists")
			}
			if competitorDTO.CompetitorID != nil {
				if err := s.tournamentQuerier.VerifyCompetitorExists(ctx, *competitorDTO.CompetitorID); err != nil {
					return customerrors.HandleErrMsg(err, "tournament", "error competitor doesn't exists")
				}
			}

			// Get the competitor who is in a tournament match.
			competitorForRemoveMatches, err := s.tournamentQuerier.GetCompetitorIDByMatchAndPosition(ctx, competitorDTO.MatchID, competitorDTO.Position)
			if err != nil {
				return customerrors.HandleErrMsg(err, "tournament", "error when adding competitors in groups")
			}

			// sacar los matches de competitorStats y cambiar las stast
			if err := s.tournamentQuerier.RemoveMultipleCompetitorStatsMatches(ctx, []string{competitorForRemoveMatches}, []string{competitorDTO.MatchID}); err != nil {
				return customerrors.HandleErrMsg(err, "tournament", "error when adding competitors in groups")
			}

			if err := s.tournamentQuerier.UpdateCompetitorMatch(ctx, competitorDTO.MatchID, competitorDTO); err != nil {
				return customerrors.HandleErrMsg(err, "tournament", "error when updating competitorMatch")
			}

		}

		courtAvailability, tournamentAvailabilities, timetablesNotAvailables, err := s.getCompleteAvailabilityInTournament(ctx, tournamentID)
		if err != nil {
			return err
		}

		competitorsInMatchMap := make(map[string][]string)
		for _, competitorsDTO := range competitorDTOs {
			// Unir los competidores con el mismo match}
			if competitorsDTO.CompetitorID != nil {
				competitorsInMatchMap[competitorsDTO.MatchID] = append(competitorsInMatchMap[competitorsDTO.MatchID], *competitorsDTO.CompetitorID)
			}
		}

		if err := s.updateMatchesDates(ctx, competitorsInMatchMap, courtAvailability, tournamentAvailabilities, timetablesNotAvailables); err != nil {
			return err
		}

		// if err := s.createMatchChats(ctx, competitorsInMatchMap, userID); err != nil {
		// 	return err
		// }

		return nil
	})
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when creating tournament location")
	}

	return nil
}

func (s *TournamentService) getCompleteAvailabilityInTournament(ctx context.Context, tournamentID string) (*dto.TournamentAvailabilityDTO, []*models.GetDailyAvailabilityByIDDTORes, []time.Time, error) {
	courtAvailability, err := s.tournamentQuerier.GetTournamentAvailavility(ctx, tournamentID)
	if err != nil {
		return nil, nil, nil, customerrors.HandleErrMsg(err, "tournament", "error when updating competitorMatches")
	}

	tournamentAvailabilities, err := s.tournamentQuerier.GetAvailabilityByTournamentID(ctx, tournamentID)
	if err != nil {
		return nil, nil, nil, customerrors.HandleErrMsg(err, "tournament", "error when updating competitorMatches")
	}

	timetablesNotAvailables, err := s.tournamentQuerier.GetAllDatesMatchesFromTournament(ctx, tournamentID)
	if err != nil {
		return nil, nil, nil, customerrors.HandleErrMsg(err, "tournament", "error when getting all dates matches from tournament")
	}

	return courtAvailability, tournamentAvailabilities, timetablesNotAvailables, nil
}
