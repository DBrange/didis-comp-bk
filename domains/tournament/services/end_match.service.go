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

func (s *TournamentService) EndMatch(ctx context.Context, match *dto.EndMatchDTOReq, secondOportunity bool) error {
	if err := s.verifyCompetitorsMatchCoincidence(ctx, match); err != nil {
		return err
	}
	if err := s.tournamentQueryer.VerifyMatchAndRoundCoincidence(ctx, match.MatchID, match.RoundID, match.Round); err != nil {
		return err
	}

	matchPosition, err := s.tournamentQueryer.GetMatchPosition(ctx, match.MatchID)
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when verifying match position")
	}

	matchesQuantity, err := s.tournamentQueryer.GetRoundQuantityMatches(ctx, match.RoundID)
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when verifying match position")
	}

	err = s.tournamentQueryer.WithTransaction(ctx, func(sessCtx mongo.SessionContext) error {
		// Set winner in match
		if err := s.tournamentQueryer.SetWinnerInMatch(ctx, match.MatchID, match.WinnerCompetitorID, match.Result); err != nil {
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

		matchesNumberForNextRound := matchesQuantity / 2

		nextRoundID, err := s.getNextRoundID(ctx, match.TournamentID, match.Round)
		if err != nil {
			return err
		}

		// Verify if next round have matches, if dont, creating
		exists, err := s.tournamentQueryer.VerifyMatchesInRoundExits(ctx, nextRoundID)
		if err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when creating matches in round")
		}
		fmt.Println(3)
		if !exists {
			if err := s.createMatchesInRound(ctx, matchesNumberForNextRound, nextRoundID, match.TournamentID, match.WinnerCompetitorID, match.Sport, secondOportunity); err != nil {
				return customerrors.HandleErrMsg(err, "tournament", "error when creating matches in round")
			}
		}
		fmt.Println(4)

		// Calculate which match and position the competitor corresponds to
		nextMatchPosition, nextPosition := s.calculatePositionsForNextRound(matchPosition)

		nextMatchID, err := s.tournamentQueryer.FindMatchID(ctx, nextMatchPosition, nextRoundID)
		if err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error not found matchID")
		}

		// In base of the number of the match ended, agregate competitor winner in next match
		if err := s.updateCompetitorMatch(ctx, nextMatchID, match.WinnerCompetitorID, nextPosition); err != nil {
			return nil
		}
		fmt.Println(6)

		return nil
	})
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when ended match")
	}

	fmt.Println(7)
	return nil
}

func (s *TournamentService) createMatchesInRound(
	ctx context.Context,
	matchesNumber int,
	roundID string,
	tournamentID,
	competitorID string,
	sport models.SPORT,
	// round models.ROUND,
	secondOportunity bool,
) error {
	for i := 0; i < int(matchesNumber); i++ {
		matchDTO := &dto.CreateMatchDTOReq{
			Sport:        sport,
			TournamentID: tournamentID,
			RoundID:      roundID,
			Position:     i + 1,
		}

		matchID, err := s.tournamentQueryer.CreateMatch(ctx, matchDTO)
		if err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when creating matches")
		}

		if err := s.CreateCompetitorMatches(ctx, matchID); err != nil {
			return err
		}


		if secondOportunity {
			// if err := s.tournamentQueryer.AddMatchInSecondOportunity(ctx, tournamentID, matchID); err != nil {
			// 	return err
			// }
		} else {
			if err := s.tournamentQueryer.AddMatchInTournament(ctx, tournamentID, matchID); err != nil {
				return err
			}
		}

		if err := s.tournamentQueryer.AddMatchInCompetitorStats(ctx, competitorID, matchID); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when updating competitor stats")
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
	competitorMatchDTO := &dto.UpdateCompetitorMatchDTOReq{CompetitorID: &competitorID, Position: &nextPosition}

	if err := s.tournamentQueryer.UpdateCompetitorMatch(ctx, nextMatchID, competitorMatchDTO); err != nil {
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

		if err := s.tournamentQueryer.UpdateCompetitorStats(ctx, competitorID, winner); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when updating competitor stats")
		}
	}

	return nil
}

func (s *TournamentService) verifyCompetitorsMatchCoincidence(ctx context.Context, match *dto.EndMatchDTOReq) error {
	if err := s.tournamentQueryer.VerifyCompetitorsMatch(ctx, match.MatchID, match.WinnerCompetitorID); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when verifying coincidence in the competitors match")
	}

	if err := s.tournamentQueryer.VerifyCompetitorsMatch(ctx, match.MatchID, match.LosserCompetitorID); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when verifying coincidence in the competitors match")
	}

	return nil
}

func (s *TournamentService) getNextRoundID(ctx context.Context, tournamentID string, round models.ROUND) (string, error) {
	nextRound, err := s.getNextRound(round)
	if err != nil {
		return "", err
	}

	nextRoundID, err := s.tournamentQueryer.GetRoundID(ctx, tournamentID, nextRound)
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
