package services

import (
	"context"
	"sort"

	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *TournamentService) EndTournament(ctx context.Context, tournamentID, doubleElimID string) error {
	err := s.tournamentQuerier.WithTransaction(ctx, func(sessCtx mongo.SessionContext) error {
		// Verify if tournament is already finished
		if err := s.tournamentQuerier.VerifyTournamentsAlreadyFinished(sessCtx, tournamentID); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when updating tournament info")
		}

		// Get champion competitor ID
		championCompetitorID, err := s.tournamentQuerier.GetCompetitorChampion(sessCtx, tournamentID)
		if err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when updating tournament info")
		}

		var doubleElimChampionCompetitorID string
		var doubleElimInfo *dto.GetDoubleElimInfoToFinaliseItDTORes

		if doubleElimID != "" {
			// Get champion competitor ID in double elimination of this tournament and double elimination info
			deChampionCompetitorID, deInfo, err := s.doubleElinationExists(sessCtx, doubleElimID)
			if err != nil {
				return err
			}

			doubleElimChampionCompetitorID = deChampionCompetitorID
			doubleElimInfo = deInfo
		}

		// Update finish_date
		if err := s.tournamentQuerier.UpdateTournamentFinishDate(sessCtx, tournamentID); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when updating tournament info")
		}

		// Add tournaments_won in competitor_stats
		if err := s.tournamentQuerier.AddTournamentWonInCompetitorStats(sessCtx, championCompetitorID, tournamentID); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when adding tournament won")
		}

		// Get rounds with his corresponding competitors
		roundDTOs, err := s.tournamentQuerier.GetRoundsWithCompetitors(sessCtx, tournamentID)
		if err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when points and/or money")
		}

		// Get tournament info
		tournamentInfo, err := s.tournamentQuerier.GetTournamentInfoToFinaliseIt(sessCtx, tournamentID)
		if err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when points and/or money")
		}

		prizeAndPoints, err := s.setPrizeAndPointsByRound(
			sessCtx,
			roundDTOs,
			tournamentInfo,
			doubleElimInfo,
			championCompetitorID,
			doubleElimChampionCompetitorID,
			doubleElimID,
		)
		if err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when points and/or money")
		}

		// Placing prize on each competitor
		for _, pm := range prizeAndPoints {
			if err := s.tournamentQuerier.AddPrizeInMultipleCompetitorStats(sessCtx, pm.CompetitorIDs, pm.TotalPrize); err != nil {
				return customerrors.HandleErrMsg(err, "tournament", "error when adding prize in competitor stats")
			}
		}

		// If tournament is not associated with any category, the following steps are not necessary
		if tournamentInfo.CategoryID != "" {
			return s.EndTournamentWithCategory(sessCtx, tournamentID, tournamentInfo.CategoryID, championCompetitorID, prizeAndPoints)
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
	competitorIDs, err := s.tournamentQuerier.GetTournamentCompetitorIDs(ctx, tournamentID)
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when getting competitorIDs")
	}

	// Get competitors
	competitorIDsOutCategory, err := s.tournamentQuerier.GetCompetitorsOutCategory(ctx, categoryID, competitorIDs)
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when finding competitors out of category")
	}

	// Add competitors
	if err := s.addCompetitorsInCategory(ctx, competitorIDsOutCategory, categoryID); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when adding points in categoryRegistration")
	}

	// Placing points on each competitor
	for _, pm := range prizeAndPoints {
		if err := s.tournamentQuerier.AddPointsInMultipleCategoryRegistration(ctx, categoryID, pm.CompetitorIDs, pm.Points); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when adding points in categoryRegistration")
		}
	}

	// reorganizar el ranking de la category en base a los nuevos puntos de los competidores
	// Si hay cambios en el ranking de algun competidor, agregarlos al slice de registered_positions (numero y hora)
	if err := s.updateCategoryRanking(ctx, categoryID); err != nil {
		return err
	}

	return nil
}

func (s *TournamentService) addCompetitorsInCategory(ctx context.Context, competitorIDsOutCategory []string, categoryID string) error {
	for _, competitorID := range competitorIDsOutCategory {
		// agregar a esos competidores que estan dentro de category
		categoryRegistrationDTO := &dto.CreateCategoryRegistrationDTOReq{
			CompetitorID: competitorID,
			CategoryID:   categoryID,
		}

		if err := s.tournamentQuerier.CreateCategoryRegistration(ctx, categoryRegistrationDTO); err != nil {
			return customerrors.HandleErrMsg(err, "category", "error when creating categoryRegistration")
		}

		if err := s.tournamentQuerier.IncrementTotalParticipants(ctx, categoryID); err != nil {
			return customerrors.HandleErrMsg(err, "category", "error when increment total participants")
		}
	}

	return nil
}

func (s *TournamentService) updateCategoryRanking(ctx context.Context, categoryID string) error {
	rankingSorted, err := s.tournamentQuerier.GetCategoryRegistrationSortedByPoints(ctx, categoryID)
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when getting categoryRegistration serted")
	}

	if err := s.tournamentQuerier.UpdateCategoryRegistrationCurrentPosition(ctx, categoryID, rankingSorted); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when updating ranking after the end of tournament")
	}

	return nil
}

func (s *TournamentService) setPrizeAndPoints(
	roundDTOs []*dto.GetRoundWithCompetitorsDTORes,
	prizeAndPointsChampions []*dto.GetRoundWithCompetitorsDTORes,
) []*dto.GetRoundWithCompetitorsDTORes {
	// Order roundDTOs from smallest to largest based on length of CompetitorIDs
	sort.Slice(roundDTOs, func(i, j int) bool {
		return len(roundDTOs[i].CompetitorIDs) < len(roundDTOs[j].CompetitorIDs)
	})

	// Add champion/s prize and points in slice
	newRoundDTOs := append(prizeAndPointsChampions, roundDTOs...)

	// We use this slice as a buffer, to know which competitors are already selected with their corresponding prize and points,
	// so that they only have the prize and points of the round that corresponds to them.
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

func (s *TournamentService) setPrizeAndPointsByRound(
	ctx context.Context,
	roundDTOs []*dto.GetRoundWithCompetitorsDTORes,
	tournamentInfo *dto.GetTournamentInfoToFinaliseItDTORes,
	doubleElimInfo *dto.GetDoubleElimInfoToFinaliseItDTORes,
	ChampionCompetitorID,
	doubleElimChampionCompetitorID,
	doubleElimID string,
) ([]*dto.GetRoundWithCompetitorsDTORes, error) {

	// Create champion
	prizeAndPointsChampion := &dto.GetRoundWithCompetitorsDTORes{
		ID:            "",
		TotalPrize:    tournamentInfo.TotalPrize,
		Points:        tournamentInfo.Points,
		CompetitorIDs: []string{ChampionCompetitorID},
	}

	// Convert to slice to be able to work
	prizeAndPointsChampionsSlice := []*dto.GetRoundWithCompetitorsDTORes{prizeAndPointsChampion}

	// In case we havent a double_elimination
	if doubleElimID == "" {
		prizeAndPoints := s.setPrizeAndPoints(roundDTOs, prizeAndPointsChampionsSlice)
		return prizeAndPoints, nil
	}

	// Create double elimination champion
	prizeAndPointsDoubleElimChampion := &dto.GetRoundWithCompetitorsDTORes{
		ID:            "",
		TotalPrize:    doubleElimInfo.TotalPrize,
		Points:        doubleElimInfo.Points,
		CompetitorIDs: []string{doubleElimChampionCompetitorID},
	}

	// Add into slice the other winner
	prizeAndPointsChampionsSlice = append(prizeAndPointsChampionsSlice, prizeAndPointsDoubleElimChampion)

	// In case we have a double_elimination
	return s.setPrizeAndPointsWithDoubleElim(ctx, roundDTOs, prizeAndPointsChampionsSlice, doubleElimID)
}

func (s *TournamentService) setPrizeAndPointsWithDoubleElim(
	ctx context.Context,
	roundDTOs []*dto.GetRoundWithCompetitorsDTORes,
	prizeAndPointsChampions []*dto.GetRoundWithCompetitorsDTORes,
	doubleElimID string,
) ([]*dto.GetRoundWithCompetitorsDTORes, error) {
	roundsInDoubleElim, err := s.tournamentQuerier.GetAllDoubleElimRoundIDs(ctx, doubleElimID)
	if err != nil {
		return nil, err
	}

	mainBracket := make([]*dto.GetRoundWithCompetitorsDTORes, 0, len(roundDTOs)-len(roundsInDoubleElim))
	secondaryBracket := make([]*dto.GetRoundWithCompetitorsDTORes, 0, len(roundsInDoubleElim))

	//  We divide the data belonging to tournament from those belonging to double elimination.
	for _, roundDTO := range roundDTOs {
		if utils.ContainsID(roundsInDoubleElim, roundDTO.ID) {
			secondaryBracket = append(secondaryBracket, roundDTO)
		} else {
			mainBracket = append(mainBracket, roundDTO)
		}
	}

	// Order the prize and points with the corresponding competitor
	mainBracketPrizeAndPoints := s.setPrizeAndPoints(mainBracket, prizeAndPointsChampions)
	secondaryBracketPrizeAndPoints := s.setPrizeAndPoints(secondaryBracket, prizeAndPointsChampions)

	mainBracketPrizeAndPointsWithoutFirstRound := mainBracketPrizeAndPoints[:len(mainBracketPrizeAndPoints)-1]
	secondaryBracketPrizeAndPointsWithoutCham := secondaryBracketPrizeAndPoints[2:]

	return append(mainBracketPrizeAndPointsWithoutFirstRound, secondaryBracketPrizeAndPointsWithoutCham...), nil
}

func (s *TournamentService) doubleElinationExists(ctx context.Context, doubleElimID string) (string, *dto.GetDoubleElimInfoToFinaliseItDTORes, error) {
	// Get champion competitor ID in double eliminacion of this tournament
	doubleElimChampionCompetitorID, err := s.tournamentQuerier.GetDoubleElimCompetitorChampion(ctx, doubleElimID)
	if err != nil {
		return "", nil, customerrors.HandleErrMsg(err, "tournament", "error when updating tournament info")
	}

	// Get double_elimiantion info
	doubleElimInfo, err := s.tournamentQuerier.GetDoubleElimInfoToFinaliseIt(ctx, doubleElimID)
	if err != nil {
		return "", nil, customerrors.HandleErrMsg(err, "tournament", "error when points and/or money")
	}

	return doubleElimChampionCompetitorID, doubleElimInfo, nil
}
