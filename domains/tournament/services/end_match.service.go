package services

import (
	"context"
	"fmt"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/mongo"
)

// RECORDAR: cuando se haga el front, si o si es necesario que el front tenga todos los id disponibles siempre

func (s *TournamentService) EndMatch(ctx context.Context, match *dto.EndMatchDTOReq) error {
	if err := s.verifyCompetitorsMatchCoincidence(ctx, match); err != nil {
		return err
	}

	matchPosition, err := s.tournamentQuerier.GetMatchPosition(ctx, match.MatchID)
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when verifying match position")
	}

	matchesQuantity, err := s.tournamentQuerier.GetRoundQuantityMatches(ctx, match.RoundID)
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when verifying match position")
	}

	err = s.tournamentQuerier.WithTransaction(ctx, func(sessCtx mongo.SessionContext) error {
		// Set winner in match
		if err := s.tournamentQuerier.SetWinnerInMatch(ctx, match.MatchID, match.WinnerCompetitorID, match.Result); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when setting match winner")
		}

		// Update stats
		if err := s.updateCompetitorStats(ctx, match); err != nil {
			return err
		}

		// If round contains less then 2 matches (is the final), only set winner
		if matchesQuantity < 2 {
			return nil
		}

		return s.handleNextRound(ctx, match, matchesQuantity, matchPosition)
	})
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when ended match")
	}

	return nil
}

func (s *TournamentService) withoutDoubleElim(ctx context.Context, match *dto.EndMatchDTOReq, competitorID string, matchesNumberForNextRound, matchPosition int) error {
	nextRoundID, err := s.getNextRoundID(ctx, match.TournamentID, match.Round)
	if err != nil {
		return err
	}

	if err := s.addCompetitorForNextRound(ctx, match, nextRoundID, competitorID, matchesNumberForNextRound, matchPosition); err != nil {
		return err
	}

	return nil
}

func (s *TournamentService) withDoubleElim(ctx context.Context, match *dto.EndMatchDTOReq, competitorID string, matchesNumberForNextRound, matchPosition int) error {
	doubleElimID, err := s.tournamentQuerier.GetDoubleElimID(ctx, match.TournamentID)
	if err != nil {
		return err
	}

	nextRoundID, err := s.getDoubleElimNextRoundID(ctx, doubleElimID, match.Round)
	if err != nil {
		return err
	}

	if err := s.addCompetitorForNextRound(ctx, match, nextRoundID, competitorID, matchesNumberForNextRound, matchPosition); err != nil {
		return err
	}

	return nil
}

func (s *TournamentService) createMatchesInRound(
	ctx context.Context,
	matchesNumber int,
	roundID string,
	tournamentID,
	competitorID string,
	sport models.SPORT,
	doubleElimID string,
) error {
	for i := 0; i < int(matchesNumber); i++ {
		matchDTO := &dto.CreateMatchDTOReq{
			Sport:        sport,
			TournamentID: tournamentID,
			RoundID:      roundID,
			Position:     i + 1,
		}

		matchID, err := s.tournamentQuerier.CreateMatch(ctx, matchDTO)
		if err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when creating matches")
		}

		if err := s.CreateCompetitorMatches(ctx, matchID); err != nil {
			return err
		}
		if doubleElimID != "" {
			if err := s.tournamentQuerier.AddMatchInDoubleElim(ctx, doubleElimID, matchID); err != nil {
				return customerrors.HandleErrMsg(err, "tournament", "error when adding match in tournament")
			}
		} else {
			if err := s.tournamentQuerier.AddMatchInTournament(ctx, tournamentID, matchID); err != nil {
				return customerrors.HandleErrMsg(err, "tournament", "error when adding match in tournament")
			}
		}

		if err := s.tournamentQuerier.AddMatchInCompetitorStats(ctx, competitorID, matchID); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when updating competitor stats")
		}
	}

	return nil
}

func (s *TournamentService) verifyMatchesInRoundExits(
	ctx context.Context,
	matchesNumberForNextRound int,
	nextRoundID,
	tournamentID,
	winnerCompetitorID string,
	sport models.SPORT,
	doubleElimID string,
) error {
	exists, err := s.tournamentQuerier.VerifyMatchesInRoundExits(ctx, nextRoundID)
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when creating matches in round")
	}
	if !exists {
		if err := s.createMatchesInRound(ctx, matchesNumberForNextRound, nextRoundID, tournamentID, winnerCompetitorID, sport, doubleElimID); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when creating matches in round")
		}
	}

	return nil
}

func (s *TournamentService) calculatePositionsForNextRound(positionMatch int) (int, int) {
	nextMatchPosition := s.calculateNextMatchPosition(positionMatch)

	nextPosition := s.calculateNextPosition(positionMatch)

	return nextMatchPosition, nextPosition
}

func (s *TournamentService) calculateNextMatchPosition(num int) int {
	return (num + 1) / 2
}

func (s *TournamentService) calculateNextPosition(num int) int {
	return ((num - 1) % 2) + 1
}

func (s *TournamentService) updateCompetitorMatch(ctx context.Context, nextMatchID, competitorID string, nextPosition int) error {
	competitorMatchDTO := &dto.UpdateCompetitorMatchDTOReq{CompetitorID: competitorID, Position: nextPosition}

	if err := s.tournamentQuerier.UpdateCompetitorMatch(ctx, nextMatchID, competitorMatchDTO); err != nil {
		return err
	}

	return nil
}

func (s *TournamentService) updateCompetitorStats(ctx context.Context, match *dto.EndMatchDTOReq) error {
	for i := 0; i < 2; i++ {
		var competitorID string
		var winner bool

		if i == 0 {
			competitorID = match.WinnerCompetitorID
			winner = true
		} else if i == 1 {
			competitorID = match.LosserCompetitorID
			winner = false
		}

		if err := s.tournamentQuerier.UpdateCompetitorStats(ctx, competitorID, winner); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when updating competitor stats")
		}
	}

	return nil
}

func (s *TournamentService) verifyCompetitorsMatchCoincidence(ctx context.Context, match *dto.EndMatchDTOReq) error {
	if err := s.tournamentQuerier.VerifyCompetitorsMatch(ctx, match.MatchID, match.WinnerCompetitorID); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when verifying coincidence in the competitors match")
	}

	if err := s.tournamentQuerier.VerifyCompetitorsMatch(ctx, match.MatchID, match.LosserCompetitorID); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when verifying coincidence in the competitors match")
	}

	if err := s.tournamentQuerier.VerifyMatchAndRoundCoincidence(ctx, match.MatchID, match.RoundID, match.Round); err != nil {
		return err
	}

	return nil
}

func (s *TournamentService) getNextRoundID(ctx context.Context, tournamentID string, round models.ROUND) (string, error) {
	nextRound, err := s.getNextRound(round)
	if err != nil {
		return "", err
	}

	nextRoundID, err := s.tournamentQuerier.GetRoundID(ctx, tournamentID, nextRound)
	if err != nil {
		return "", customerrors.HandleErrMsg(err, "tournament", "error when getting roundID")
	}

	return nextRoundID, nil
}

func (s *TournamentService) getDoubleElimNextRoundID(ctx context.Context, doubleElimID string, round models.ROUND) (string, error) {
	nextRound, err := s.getNextRound(round)
	if err != nil {
		return "", err
	}

	nextRoundID, err := s.tournamentQuerier.GetDoubleElimRoundID(ctx, doubleElimID, nextRound)
	if err != nil {
		return "", customerrors.HandleErrMsg(err, "tournament", "error when getting roundID")
	}

	return nextRoundID, nil
}

func (s *TournamentService) getNextRound(round models.ROUND) (models.ROUND, error) {
	roundMap := map[models.ROUND]models.ROUND{
		models.ROUND_1R: models.ROUND_2R,
		models.ROUND_2R: models.ROUND_3R,
		models.ROUND_3R: models.ROUND_4R,
		models.ROUND_4R: models.ROUND_5R,
		models.ROUND_5R: models.ROUND_OF,
		models.ROUND_OF: models.ROUND_CF,
		models.ROUND_CF: models.ROUND_SF,
		models.ROUND_SF: models.ROUND_F,
	}

	for currentRound, nextRound := range roundMap {
		if currentRound == round {
			return nextRound, nil
		}
	}

	err := fmt.Errorf("error round name not exists")
	return models.ROUND(""), customerrors.HandleErrMsg(err, "tournament", "error round name not exists")
}

func (s *TournamentService) handleNextRound(ctx context.Context, match *dto.EndMatchDTOReq, matchesQuantity, matchPosition int) error {
	matchesNumberForNextRound := matchesQuantity / 2

	roundNames, err := s.tournamentQuerier.GetTournamentRoundNames(ctx, match.TournamentID)
	if err != nil {
		return err
	}

	firstRoundInTournament := s.calculateFirstRoundInTournament(roundNames)

	if match.Round == firstRoundInTournament && match.DoubleElimID != "" {
		// Add competitor to the next round in main drow and add competitor to double elimination
		return s.handleFirstRound(ctx, match, matchesNumberForNextRound, matchPosition)
	}

	if match.DoubleElimID != "" {
		// Add competitor to the next round in double elimination
		return s.withDoubleElim(ctx, match, match.WinnerCompetitorID, matchesNumberForNextRound, matchPosition)
	}

	// Add competitor to the next round in main draw
	return s.withoutDoubleElim(ctx, match, match.WinnerCompetitorID, matchesNumberForNextRound, matchPosition)
}

func (s *TournamentService) handleFirstRound(ctx context.Context, match *dto.EndMatchDTOReq, matchesNumberForNextRound, matchPosition int) error {
	if err := s.withDoubleElim(ctx, match, match.LosserCompetitorID, matchesNumberForNextRound, matchPosition); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error withDoubleElim")
	}

	if err := s.withoutDoubleElim(ctx, match, match.WinnerCompetitorID, matchesNumberForNextRound, matchPosition); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error withoutDoubleElim")
	}

	return nil
}

func (s *TournamentService) calculateFirstRoundInTournament(rounds []models.ROUND) models.ROUND {
	roundMap := map[models.ROUND]int{
		models.ROUND_1R: 1,
		models.ROUND_2R: 2,
		models.ROUND_3R: 3,
		models.ROUND_4R: 4,
		models.ROUND_5R: 5,
		models.ROUND_OF: 6,
		models.ROUND_CF: 7,
		models.ROUND_SF: 8,
		models.ROUND_F:  9,
	}

	if len(rounds) == 0 {
		return "" // O algún valor por defecto que tenga sentido en tu contexto
	}

	lowestRound := rounds[0]
	lowestValue := roundMap[lowestRound]

	for _, round := range rounds {
		if value, exists := roundMap[round]; exists && value < lowestValue {
			lowestRound = round
			lowestValue = value
		}
	}

	return lowestRound
}

func (s *TournamentService) addCompetitorForNextRound(ctx context.Context, match *dto.EndMatchDTOReq, nextRoundID, competitorID string, matchesNumberForNextRound int, matchPosition int) error {
	// Verify if next round have matches, if dont, creating
	if err := s.verifyMatchesInRoundExits(ctx, matchesNumberForNextRound, nextRoundID, match.TournamentID, match.WinnerCompetitorID, match.Sport, match.DoubleElimID); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when creating matches in round")
	}

	// Calculate which match and position the competitor corresponds to
	nextMatchPosition, nextPosition := s.calculatePositionsForNextRound(matchPosition)

	nextMatchID, err := s.tournamentQuerier.FindMatchID(ctx, nextMatchPosition, nextRoundID)
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error not found matchID")
	}

	// In base of the number of the match ended, agregate competitor winner in next match
	if err := s.updateCompetitorMatch(ctx, nextMatchID, competitorID, nextPosition); err != nil {
		return err
	}

	return nil
}
