package services

import (
	"context"
	"fmt"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	option_models "github.com/DBrange/didis-comp-bk/cmd/api/models/options/tournament"
	"github.com/DBrange/didis-comp-bk/domains/tournament/adapters/mappers"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *TournamentService) OrganizeTournament(ctx context.Context, organizeTournamentDTO *dto.OrganizeTournamentDTOReq, options *option_models.OrganizeTournamentOptions) error {
	err := s.tournamentQueryer.WithTransaction(ctx, func(sessCtx mongo.SessionContext) error {
		// Mapping info
		tournamentDTO, locationDTO, categoryID, organizerID := mappers.OrganizeTournamentMapper(organizeTournamentDTO)

		// Verifications
		if err := s.verifications(sessCtx, organizerID, categoryID); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error organizer not exists")
		}

		// Create location for tournament
		locationID, err := s.tournamentQueryer.CreateLocation(sessCtx, locationDTO)
		if err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when creating tournament location")
		}

		// Create tournament
		tournamentID, err := s.tournamentQueryer.CreateTournament(
			sessCtx,
			tournamentDTO,
			locationID,
			options,
			categoryID,
			organizerID,
		)
		if err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when creating tournament")
		}

		// If a category ID is available, add this tournament at the category
		if categoryID != nil {
			if err := s.tournamentQueryer.AddTournamentInCategory(sessCtx, *categoryID, tournamentID); err != nil {
				return customerrors.HandleErrMsg(err, "tournament", "error when adding tournament to category")
			}
		}

		// Add pots and/or groups in tournament
		if options.QuantityGroups > 0 || options.QuantityPots > 0 {
			if err = s.tournamentWithPotsGroups(sessCtx, options, tournamentID); err != nil {
				return err
			}
		}

		// Without pots and groups, add aumatically the rounds and firts matches into tournament
		if options.QuantityGroups <= 0 && options.QuantityPots <= 0 {
			if err := s.onlyBrackets(sessCtx, tournamentID, tournamentDTO, options); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when creating tournament location")
	}

	return nil
}

// Add pots and/or groups in tournament
func (s *TournamentService) tournamentWithPotsGroups(
	sessCtx mongo.SessionContext,
	options *option_models.OrganizeTournamentOptions,
	tournamentID string,
) error {
	var tournamentOptions dto.UpdateTournamentOptionsDTOReq

	// Update the tournament slice of pots and/or groups
	err := s.updatePotsGroupsSlice(sessCtx, options, &tournamentOptions, tournamentID)
	if err != nil {
		return err
	}

	// Update tournament
	if err = s.tournamentQueryer.UpdateTournamentRelations(sessCtx, tournamentID, &tournamentOptions, true); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when updating tournament")
	}

	return nil
}

// Update the tournament slice of pots and/or groups
func (s *TournamentService) updatePotsGroupsSlice(
	sessCtx mongo.SessionContext,
	options *option_models.OrganizeTournamentOptions,
	tournamentOptions *dto.UpdateTournamentOptionsDTOReq,
	tournamentID string,
) error {
	// Create Pots slice
	if options.QuantityPots > 0 {
		tournamentOptions.Pots = &[]string{}
		err := s.createPotsGroupsSlice(sessCtx, tournamentID, options.QuantityPots, "pot", tournamentOptions.Pots, s.tournamentQueryer.CreatePot)
		if err != nil {
			return err
		}
	}

	//Create groups slice
	if options.QuantityGroups > 0 {
		tournamentOptions.Groups = &[]string{}
		err := s.createPotsGroupsSlice(sessCtx, tournamentID, options.QuantityGroups, "group", tournamentOptions.Groups, s.tournamentQueryer.CreateTournamentGroup)
		if err != nil {
			return err
		}
		roundID, err := s.createRoundsForGroup(sessCtx, tournamentID)
		if err != nil {
			return err
		}

		roundSlice := []string{roundID}
		if err := s.updateTournament(sessCtx, tournamentID, nil, &roundSlice, nil); err != nil {
			return err
		}

	}

	// Create double elimination
	if options.DoubleElimination {
		doubleEliminationID, err := s.tournamentQueryer.CreateDoubleEliminationEmpty(sessCtx)
		if err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when creating double elimination")
		}

		tournamentOptions.DoubleEliminationID = &doubleEliminationID
	}

	return nil
}

type FnCreate func(ctx context.Context, tournamentID string, postion int) (string, error)

// Create slice of pots and/or groups
func (s *TournamentService) createPotsGroupsSlice(
	sessCtx mongo.SessionContext,
	tournamentID string,
	quantity int,
	name string,
	dest *[]string,
	fnCreate FnCreate,
) error {
	IDs := make([]string, quantity)
	for i := 0; i < quantity; i++ {
		vID, err := fnCreate(sessCtx, tournamentID, i+1) // valueID
		if err != nil {
			errMsg := fmt.Sprintf("error when creating %sID", name)
			return customerrors.HandleErrMsg(err, "tournament", errMsg)
		}

		IDs[i] = vID
	}

	*dest = IDs

	return nil
}

func (s *TournamentService) calculateRoundsNumber(matchesNumber models.TOURNAMENT_MATCHES_CAPACITY) (int, error) {
	// Map of all options in enum TOURNAMENT_CAPACITY
	quantityOpts := models.ALL_TOURNAMENT_MATCHES_CAPACITY
	// If number of matches is valid, the corresponding number of rounds required will be sent
	for i, quantity := range quantityOpts {
		if quantity == matchesNumber {
			return i + 1, nil
		}
	}

	err := fmt.Errorf("invalid total competitors number")
	return 0, customerrors.HandleErrMsg(err, "tournament", "error when calculating rounds number")
}

func (s *TournamentService) calculateMatchesNumber(maxCapacity models.TOURNAMENT_CAPACITY, competitorType models.COMPETITOR_TYPE) (int, error) {
	quantityOpts := models.ALL_TOURNAMENT_CAPACITY

	competitorTypeOpts := map[models.COMPETITOR_TYPE]int{
		models.COMPETITOR_TYPE_SINGLE: 1,
		models.COMPETITOR_TYPE_DOUBLE: 2,
		models.COMPETITOR_TYPE_TEAM:   1,
	}

	// Calculate based on competitors, the number of the matches to be played, and create them
	//if competitonType is "D" halving the number in matches (doubles: 4 competitor for match), else calculate normal
	// If number of total competitors is valid, the corresponding number of matches required will be sent
	for _, quantity := range quantityOpts {
		if quantity == maxCapacity {
			matchesNumber := int(quantity) / 2 / int(competitorTypeOpts[competitorType])
			if competitorType == models.COMPETITOR_TYPE_DOUBLE {
				// Verify if the number of competitors is odd
				if matchesNumber%2 != 0 {
					err := fmt.Errorf("invalid total competitors number")
					return 0, customerrors.HandleErrMsg(err, "tournament", "error when calculating matches number")
				}
			}
			return matchesNumber, nil
		}
	}

	err := fmt.Errorf("invalid total competitors number")
	return 0, customerrors.HandleErrMsg(err, "tournament", "error when calculating matches number")
}

func (s *TournamentService) calculateRoundName(i int) models.ROUND {
	nameMap := map[int]models.ROUND{
		0: models.ROUND_F,
		1: models.ROUND_SF,
		2: models.ROUND_CF,
		3: models.ROUND_OF,
		4: models.ROUND_1R,
		5: models.ROUND_2R,
		6: models.ROUND_3R,
		7: models.ROUND_4R,
		8: models.ROUND_5R,
	}

	return nameMap[i]
}

func (s *TournamentService) calculateRoundInMatch(roundsCreated []string) string {
	maxIndex := 4 // max index
	if len(roundsCreated) == 0 {
		err := fmt.Errorf("error when calculare round in match, value 0")
		customerrors.HandleErrMsg(err, "tournament", "error when calculare round in match, value 0")
	}

	// Make sure not to exeed max index
	index := len(roundsCreated) - 1
	if index > maxIndex {
		index = maxIndex
	}

	// this id the logic:
	//  matches    |    round
	//  	 1              F
	//     2              SF
	//     4              CF
	//     8              OF
	//     16             1R
	//     32             2R
	//     64             3R
	//     128            4R
	//     256            5R

	return roundsCreated[index]
}

func (s *TournamentService) createRounds(ctx context.Context, roundsNumber int, tournamentID string) ([]string, error) {
	roundsCreated := make([]string, roundsNumber)
	for i := 0; i < roundsNumber; i++ {
		roundDTO := &dto.CreateRoundDTOReq{
			TournamentID: tournamentID,
			Name:         s.calculateRoundName(i),
		}

		roundID, err := s.tournamentQueryer.CreateRound(ctx, roundDTO)
		if err != nil {
			return []string{}, customerrors.HandleErrMsg(err, "tournament", "error when creating rounds")
		}

		roundsCreated[i] = roundID
	}

	return roundsCreated, nil
}
func (s *TournamentService) createRoundsForGroup(ctx context.Context, tournamentID string) (string, error) {
	roundDTO := &dto.CreateRoundDTOReq{
		TournamentID: tournamentID,
		Name:         models.ROUND_GROUP,
	}
	roundID, err := s.tournamentQueryer.CreateRound(ctx, roundDTO)
	if err != nil {
		return "", customerrors.HandleErrMsg(err, "tournament", "error when creating rounds")
	}

	return roundID, nil
}

func (s *TournamentService) createMatches(ctx context.Context, matchesNumber int, roundsCreated []string, tournamentID string, sport models.SPORT) ([]string, error) {
	matchesCreated := make([]string, matchesNumber)
	for i := 0; i < int(matchesNumber); i++ {
		matchDTO := &dto.CreateMatchDTOReq{
			Sport:        sport,
			TournamentID: tournamentID,
			RoundID:      s.calculateRoundInMatch(roundsCreated),
			Position:     i + 1,
		}

		matchID, err := s.tournamentQueryer.CreateMatch(ctx, matchDTO)
		if err != nil {
			return []string{}, customerrors.HandleErrMsg(err, "tournament", "error when creating matches")
		}

		if err := s.CreateCompetitorMatches(ctx, matchID); err != nil {
			return []string{}, err
		}

		matchesCreated[i] = matchID
	}

	return matchesCreated, nil
}

func (s *TournamentService) updateTournament(ctx context.Context, tournamentID string, doubleEliminationID *string, roundsCreated *[]string, matchesCreated *[]string) error {
	tournamentOptsDTO := &dto.UpdateTournamentOptionsDTOReq{
		DoubleEliminationID: doubleEliminationID,
		Rounds:              roundsCreated,
		Matches:             matchesCreated,
	}

	// Update tournament with the new matches, rounds and doubleElimination(if it is true)
	if err := s.tournamentQueryer.UpdateTournamentRelations(ctx, tournamentID, tournamentOptsDTO, true); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when updating tournament")
	}

	return nil
}

func (s *TournamentService) verifications(ctx context.Context, organizerID string, categoryID *string) error {
	// Verify if organizer exists
	if err := s.tournamentQueryer.VerifyOrganizerExists(ctx, organizerID); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error organizer not exists")
	}

	// Verify if category exists
	if categoryID != nil {
		if err := s.tournamentQueryer.VerifyCategoryExists(ctx, *categoryID); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error category not exists")
		}
	}

	return nil
}

// Without pots and groups, add aumatically the rounds and firts matche into tournaments
func (s *TournamentService) onlyBrackets(ctx context.Context, tournamentID string, tournamentDTO *dto.CreateTournamentDTOReq, options *option_models.OrganizeTournamentOptions) error {
	if options.QuantityGroups <= 0 && options.QuantityPots <= 0 {
		matchesNumber, roundsNumber, err := s.calculateQuantityMatchesRounds(tournamentDTO.MaxCapacity, tournamentDTO.CompetitorType)
		if err != nil {
			return err
		}

		matchesCreated, roundsCreated, err := s.createMatchesRounds(ctx, tournamentID, tournamentDTO.Sport, roundsNumber, matchesNumber)
		if err != nil {
			return err
		}

		var doubleEliminationID string

		// if DoubleElimination is true, create and also calculate their matches and rounds, always halving the number of main draw matches
		if options.DoubleElimination {
			doubleElimID, err := s.createDoubleElimination(ctx, tournamentID, tournamentDTO, int(matchesNumber), roundsNumber)
			if err != nil {
				return err
			}

			doubleEliminationID = doubleElimID
		}

		// Update tournament with the new matches, rounds and doubleElimination(if it is true)
		if err := s.updateTournament(ctx, tournamentID, &doubleEliminationID, &roundsCreated, &matchesCreated); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when updating tournament")
		}
	}

	return nil
}

func (s *TournamentService) createDoubleElimination(ctx context.Context, tournamentID string, tournamentDTO *dto.CreateTournamentDTOReq, matchesNumber, roundsNumber int) (string, error) {
	matchesNumberde := matchesNumber / 2
	roundsNumberde := roundsNumber - 1

	matchesCreated, roundsCreated, err := s.createMatchesRounds(ctx, tournamentID, tournamentDTO.Sport, roundsNumberde, matchesNumberde)
	if err != nil {
		return "", err
	}

	doubleEliminationDTO := &dto.CreateDoubleEliminationDTOReq{
		Matches: matchesCreated,
		Rounds:  roundsCreated,
	}

	doubleEliminationID, err := s.tournamentQueryer.CreateDoubleElimination(ctx, doubleEliminationDTO)
	if err != nil {
		return "", customerrors.HandleErrMsg(err, "tournament", "error when creating double elimination")
	}

	return doubleEliminationID, nil
}

func (s *TournamentService) calculateQuantityMatchesRounds(maxCapacity models.TOURNAMENT_CAPACITY, competitorsType models.COMPETITOR_TYPE) (int, int, error) {
	totalMatches, err := s.calculateMatchesNumber(maxCapacity, competitorsType)
	if err != nil {
		return 0, 0, err
	}

	matchesNumber := models.TOURNAMENT_MATCHES_CAPACITY(totalMatches)

	// Calculate the number of rounds and create them
	roundsNumber, err := s.calculateRoundsNumber(matchesNumber)
	if err != nil {
		return 0, 0, err
	}

	return totalMatches, roundsNumber, nil
}

func (s *TournamentService) createMatchesRounds(ctx context.Context, tournamentID string, sport models.SPORT, roundsNumber, matchesNumber int) ([]string, []string, error) {
	roundsCreated, err := s.createRounds(ctx, roundsNumber, tournamentID)
	if err != nil {
		return []string{}, []string{}, err
	}

	matchesCreated, err := s.createMatches(ctx, int(matchesNumber), roundsCreated, tournamentID, sport)
	if err != nil {
		return []string{}, []string{}, err
	}

	return matchesCreated, roundsCreated, err
}

func (s *TournamentService) CreateCompetitorMatches(ctx context.Context, matchID string) error {
	for i := 0; i < 2; i++ {
		competitorMatch := &dto.CreateCompetitorMatchDTOReq{
			Position:     i + 1,
			CompetitorID: nil,
			MatchID:      matchID,
		}

		// create competitor_match
		if err := s.tournamentQueryer.CreateCompetitorMatch(ctx, competitorMatch); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when creating competitorMatches")
		}
	}

	return nil
}
