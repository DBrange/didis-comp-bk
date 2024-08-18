package services

import (
	"context"
	"sort"

	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *TournamentService) EndTournament(ctx context.Context, tournamentID, competitorID string) error {
	err := s.tournamentQueryer.WithTransaction(ctx, func(sessCtx mongo.SessionContext) error {
		// Verify if tournament is already finished
		if err := s.tournamentQueryer.VerifyTournamentsAlreadyFinished(ctx, tournamentID); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when updating tournament info")
		}

		// Get tournament info
		tournamentInfo, err := s.tournamentQueryer.GetTournamentInfoToFinaliseIt(ctx, tournamentID)
		if err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when points and/or money")
		}

		// Update finish_date
		if err := s.tournamentQueryer.UpdateTournamentFinishDate(ctx, tournamentID); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when updating tournament info")
		}

		// Add tournaments_won in competitor_stats
		if err := s.tournamentQueryer.AddTournamentWonInCompetitorStats(ctx, competitorID, tournamentID); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when adding tournament won")
		}

		// Get rounds with his corresponding competitors
		roundDTOs, err := s.tournamentQueryer.GetRoundsWithCompetitors(ctx, tournamentID)
		if err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when points and/or money")
		}

		// Order the prize and points with the corresponding competitor
		prizeAndPoints := s.setPrizeAndPoints(roundDTOs, tournamentInfo.TotalPrize, tournamentInfo.Points, competitorID)

		// Placing prize on each competitor
		for _, pm := range prizeAndPoints {
			if err := s.tournamentQueryer.AddPrizeInMultipleCompetitorStats(ctx, pm.CompetitorIDs, pm.TotalPrize); err != nil {
				return customerrors.HandleErrMsg(err, "tournament", "error when adding prize in competitor stats")
			}
		}

		// If tournament is not associated with any category, the following steps are not necessary
		if competitorID != "" {
			return s.EndTournamentWithCategory(ctx, tournamentID, tournamentInfo.CategoryID, competitorID, prizeAndPoints)
		}

		return nil
	})
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when ended match")
	}

	return nil
}

func (s *TournamentService) EndTournamentWithCategory(ctx context.Context, tournamentID, categoryID, competitorID string, prizeAndPoints []*dto.GetRoundWithCompetitorsDTORes) error {
	// Get competitors in competitors_registration
	competitorIDs, err := s.tournamentQueryer.GetTournamentCompetitorIDs(ctx, tournamentID)
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when getting competitorIDs")
	}

	// Get competitors
	competitorIDsOutCategory, err := s.tournamentQueryer.GetCompetitorsOutCategory(ctx, categoryID, competitorIDs)
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when finding competitors out of category")
	}

	for _, competitorID := range competitorIDsOutCategory {
		// agregar a esos competidores que estan dentro de category
		categoryRegistrationDTO := &dto.CreateCategoryRegistrationDTOReq{
			CompetitorID: competitorID,
			CategoryID:   categoryID,
		}

		if err := s.tournamentQueryer.CreateCategoryRegistration(ctx, categoryRegistrationDTO); err != nil {
			return customerrors.HandleErrMsg(err, "category", "error when creating categoryRegistration")
		}

		if err := s.tournamentQueryer.IncrementTotalParticipants(ctx, categoryID); err != nil {
			return customerrors.HandleErrMsg(err, "category", "error when increment total participants")
		}
	}

	// reorganizar el ranking de la category en base a los nuevos puntos de los competidores
	// Si hay cambios en el ranking de algun competidor, agregarlos al slice de registered_positions (numero y hora)
	if err := s.updateCategoryRanking(ctx, categoryID); err != nil {
		return err
	}

	// Placing points on each competitor
	for _, pm := range prizeAndPoints {
		if err := s.tournamentQueryer.AddPointsInMultipleCategoryRegistration(ctx, categoryID, pm.CompetitorIDs, pm.Points); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when adding points in categoryRegistration")
		}
	}

	return nil
}

func (s *TournamentService) updateCategoryRanking(ctx context.Context, categoryID string) error {
	rankingSorted, err := s.tournamentQueryer.GetCategoryRegistrationSortedByPoints(ctx, categoryID)
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when getting categoryRegistration serted")
	}

	if err := s.tournamentQueryer.UpdateCategoryRegistrationCurrentPosition(ctx, categoryID, rankingSorted); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when updating ranking after the end of tournament")
	}

	return nil
}

func (s *TournamentService) setPrizeAndPoints(roundDTOs []*dto.GetRoundWithCompetitorsDTORes, total_prize float64, points int, competitorWinnerID string) []*dto.GetRoundWithCompetitorsDTORes {
	// Ordenar los roundDTOs de menor a mayor basado en la longitud de CompetitorIDs
	sort.Slice(roundDTOs, func(i, j int) bool {
		return len(roundDTOs[i].CompetitorIDs) < len(roundDTOs[j].CompetitorIDs)
	})

	prizeAndPointsWinner := &dto.GetRoundWithCompetitorsDTORes{
		ID:            "",
		TotalPrize:    total_prize,
		Points:        points,
		CompetitorIDs: []string{competitorWinnerID},
	}
	prizeAndPointsWinnerSlice := []*dto.GetRoundWithCompetitorsDTORes{prizeAndPointsWinner}
	newRoundDTOs := append(prizeAndPointsWinnerSlice, roundDTOs...)

	usedCompetitors := make(map[string]bool)

	for i := 0; i < len(newRoundDTOs); i++ {
		newCompetitors := []string{}
		for _, comp := range newRoundDTOs[i].CompetitorIDs {
			if !usedCompetitors[comp] {
				newCompetitors = append(newCompetitors, comp)
				usedCompetitors[comp] = true
			}
		}
		newRoundDTOs[i].CompetitorIDs = newCompetitors
	}

	return newRoundDTOs
}
