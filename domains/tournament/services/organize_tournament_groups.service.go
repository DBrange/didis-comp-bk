package services

import (
	"context"
	"math/rand"
	"sort"
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *TournamentService) OrganizeTournamentGroups(ctx context.Context, tournamentID, roundID string, competitorDTOs []*dto.AddCompetitorsToTournamentGroupsDTOReq, sport models.SPORT, orderType, top int) error {
	categoryID, err := s.tournamentQuerier.GetCategoryIDOfTournament(ctx, tournamentID)
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when getting tournament competitors")
	}

	err = s.tournamentQuerier.WithTransaction(ctx, func(sessCtx mongo.SessionContext) error {
		competitorsDTO, err := s.tournamentQuerier.GetCompetitorsInTournament(ctx, tournamentID, categoryID, "", 0, true)
		if err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when getting tournament competitors")
		}

		groupIDs, err := s.tournamentQuerier.GetTournamentGroupsIDs(ctx, tournamentID)
		if err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when getting tournament competitors")
		}

		var competitors []*dto.AddCompetitorsToTournamentGroupsDTOReq

		if orderType == 0 {
			competitors = distributeCompetitorsRandom(groupIDs, competitorsDTO)
		} else {
			competitors = distributeCompetitorsByPosition(groupIDs, competitorsDTO, top)
		}

		// Add competitors on each group
		if err := s.tournamentQuerier.AddCompetitorsToTournamentGroups(sessCtx, tournamentID, competitors); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when adding competitors in groups")
		}

		// Create matches alghoritm
		if err := s.createRoundRobin(sessCtx, tournamentID, roundID, competitors, sport); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when adding competitors in groups")
		}

		return nil
	})
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when creating tournament location")
	}

	return nil
}

func distributeCompetitorsRandom(groupIDs []string, competitors []*dto.GetCompetitorsInTournamentCompetitorDTORes) []*dto.AddCompetitorsToTournamentGroupsDTOReq {
	// Create a copy of competitors to avoid modifying the original slice
	shuffledCompetitors := make([]*dto.GetCompetitorsInTournamentCompetitorDTORes, len(competitors))
	copy(shuffledCompetitors, competitors)

	// Create a new random generator with a time-based seed
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Shuffle the competitors using the new random generator
	r.Shuffle(len(shuffledCompetitors), func(i, j int) {
		shuffledCompetitors[i], shuffledCompetitors[j] = shuffledCompetitors[j], shuffledCompetitors[i]
	})

	// Create a slice to hold the result
	groups := make([]*dto.AddCompetitorsToTournamentGroupsDTOReq, len(groupIDs))

	// Calculate the number of competitors per group
	numGroups := len(groupIDs)
	chunkSize := len(shuffledCompetitors) / numGroups
	remaining := len(shuffledCompetitors) % numGroups

	// Distribute competitors among the groups
	index := 0
	for i, groupID := range groupIDs {
		// Calculate how many competitors to add to this group
		size := chunkSize
		if i < remaining {
			// Distribute remaining competitors among first groups
			size++
		}

		// Extract competitor IDs for this group
		competitorIDs := make([]string, size)
		for j := 0; j < size; j++ {
			competitorIDs[j] = shuffledCompetitors[index+j].CompetitorID
		}

		// Assign competitors slice to the group
		groups[i] = &dto.AddCompetitorsToTournamentGroupsDTOReq{
			GroupID:     groupID,
			Competitors: competitorIDs,
		}
		index += size
	}

	return groups
}

func distributeCompetitorsByPosition(
	groupIDs []string,
	competitors []*dto.GetCompetitorsInTournamentCompetitorDTORes,
	top int,
) []*dto.AddCompetitorsToTournamentGroupsDTOReq {
	// Limitar 'top' al tamaño de competidores para evitar índices fuera de rango
	if top < 0 {
		top = 0
	}
	if top > len(competitors) {
		top = len(competitors)
	}

	// Separar competidores en válidos e inválidos
	var validCompetitors []dto.GetCompetitorsInTournamentCompetitorDTORes
	var invalidCompetitors []dto.GetCompetitorsInTournamentCompetitorDTORes

	for _, competitor := range competitors {
		if competitor.CurrentPosition != nil && *competitor.CurrentPosition > 0 {
			validCompetitors = append(validCompetitors, *competitor)
		} else {
			invalidCompetitors = append(invalidCompetitors, *competitor)
		}
	}

	// Ordenar los competidores válidos por current_position
	sort.Slice(validCompetitors, func(i, j int) bool {
		return *validCompetitors[i].CurrentPosition < *validCompetitors[j].CurrentPosition
	})

	// Final competitors array
	var finalCompetitors []dto.GetCompetitorsInTournamentCompetitorDTORes

	// Si top es 0, no necesitamos mezclar y simplemente se distribuyen todos los competidores válidos
	if top == 0 {
		finalCompetitors = validCompetitors // Todos los válidos
	} else {
		// Si hay más competidores que top, mezclar los restantes después de los top
		if top < len(validCompetitors) {
			// Mezclar el resto de los competidores válidos (los que están después de los primeros `top`)
			remainingCompetitors := validCompetitors[top:]

			// Crear un nuevo generador de números aleatorios
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			r.Shuffle(len(remainingCompetitors), func(i, j int) {
				remainingCompetitors[i], remainingCompetitors[j] = remainingCompetitors[j], remainingCompetitors[i]
			})

			// Combinar los primeros `top` con los restantes mezclados
			finalCompetitors = append(validCompetitors[:top], remainingCompetitors...)
		} else {
			// Si no hay competidores restantes, solo tenemos los válidos
			finalCompetitors = validCompetitors
		}
	}

	// Añadir inválidos al final
	finalCompetitors = append(finalCompetitors, invalidCompetitors...)

	// Distribuir competidores en grupos de manera cíclica
	groups := make([]*dto.AddCompetitorsToTournamentGroupsDTOReq, len(groupIDs))
	for i := range groupIDs {
		groups[i] = &dto.AddCompetitorsToTournamentGroupsDTOReq{
			GroupID:     groupIDs[i],
			Competitors: []string{},
		}
	}

	// Distribuir cíclicamente
	for i, competitor := range finalCompetitors {
		groupIndex := i % len(groupIDs)
		groups[groupIndex].Competitors = append(groups[groupIndex].Competitors, competitor.CompetitorID)
	}

	return groups
}

func (s *TournamentService) createRoundRobin(ctx context.Context, tournamentID, roundID string, competitorDTOs []*dto.AddCompetitorsToTournamentGroupsDTOReq, sport models.SPORT) error {
	competitorsInMatchMap, err := s.createMatchesFromRoundRobin(ctx, tournamentID, roundID, competitorDTOs, sport)
	if err != nil {
		return err
	}

	courtAvailability, tournamentAvailabilities, err := s.getPartialAvailabilityInTournament(ctx, tournamentID)
	if err != nil {
		return err
	}

	timetablesNotAvailables := []time.Time{}
	if err := s.updateMatchesDates(ctx, competitorsInMatchMap, courtAvailability, tournamentAvailabilities, timetablesNotAvailables); err != nil {
		return err
	}

	return nil
}

func (s *TournamentService) createMatchesFromRoundRobin(
	ctx context.Context,
	tournamentID,
	roundID string,
	competitorDTOs []*dto.AddCompetitorsToTournamentGroupsDTOReq,
	sport models.SPORT,
) (map[string][]string, error) {
	competitorsInMatchMap := make(map[string][]string)

	for _, competitorDTO := range competitorDTOs {
		competitors := competitorDTO.Competitors

		matchesNum := s.CalculateTotalMatchesInGroup(competitors) // Número total de partidos en un sistema "todos contra todos"

		competitorMatchDTO := make([]*dto.UpdateCompetitorMatchDTOReq, matchesNum*2) // *2 porque cada partido tiene 2 competidores

		// Create matches
		matchIDs, err := s.createMatchesInGroup(ctx, matchesNum, roundID, tournamentID, sport)
		if err != nil {
			return nil, customerrors.HandleErrMsg(err, "tournament", "error when creating matches")
		}

		// Add matches in tournament
		if err := s.addMatchesInTournament(ctx, tournamentID, competitorDTO.GroupID, matchIDs); err != nil {
			return nil, customerrors.HandleErrMsg(err, "tournament", "error when adding competitors in groups")
		}

		matchStatsMap, err := s.addMatchInStats(ctx, matchIDs, competitorMatchDTO, competitors)
		if err != nil {
			return nil, err
		}

		for key, value := range matchStatsMap {
			competitorsInMatchMap[key] = append(competitorsInMatchMap[key], value...)
		}

		// Update competitorMatches with the new info
		if err := s.tournamentQuerier.UpdateMultipleCompetitorMatches(ctx, competitorMatchDTO); err != nil {
			return nil, customerrors.HandleErrMsg(err, "tournament", "error when updating competitorMatches")
		}
	}

	return competitorsInMatchMap, nil

}

func (s *TournamentService) addMatchInStats(
	ctx context.Context,
	matchIDs []string,
	competitorMatchDTO []*dto.UpdateCompetitorMatchDTOReq,
	competitors []string,
) (map[string][]string, error) {
	competitorsInMatchMap := make(map[string][]string)

	n := len(competitors)

	matchIndex := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if matchIndex >= len(matchIDs) {
				break
			}
			matchID := matchIDs[matchIndex]

			// Competitor objects
			competitorMatchDTO[matchIndex*2] = &dto.UpdateCompetitorMatchDTOReq{
				MatchID:      matchID,
				CompetitorID: &competitors[i],
				Position:     *utils.IntPtr(1),
			}
			competitorMatchDTO[matchIndex*2+1] = &dto.UpdateCompetitorMatchDTOReq{
				MatchID:      matchID,
				CompetitorID: &competitors[j],
				Position:     *utils.IntPtr(2),
			}

			// Add competitors to stats
			if err := s.tournamentQuerier.AddMatchInCompetitorStats(ctx, competitors[i], matchID); err != nil {
				return nil, customerrors.HandleErrMsg(err, "tournament", "error when adding match competitor stats")
			}
			if err := s.tournamentQuerier.AddMatchInCompetitorStats(ctx, competitors[j], matchID); err != nil {
				return nil, customerrors.HandleErrMsg(err, "tournament", "error when adding match competitor stats")
			}

			competitorsInMatchMap[matchID] = append(competitorsInMatchMap[matchID], competitors[i], competitors[j])

			matchIndex++
		}
	}

	return competitorsInMatchMap, nil
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
		matchID, err := s.tournamentQuerier.CreateMatch(ctx, matchDTO)
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
	if err := s.tournamentQuerier.VerifyMultipleGroupsInTournament(ctx, tournamentID, groupIDs); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when adding competitors in groups")
	}

	// Verify if round is on tournament
	if err := s.tournamentQuerier.VerifyRoundInTournament(ctx, roundID, tournamentID); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when adding competitors in groups")
	}

	// Verify if tournament contain all competitors
	if err := s.tournamentQuerier.VerifyMultipleCompetitorsExistsInTournament(ctx, tournamentID, competitorIDs); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when verifying competitors")
	}

	return nil
}

func (s *TournamentService) addMatchesInTournament(ctx context.Context, tournamentID, groupID string, matchIDs []string) error {
	// Add matches in tournament
	// if err := s.tournamentQuerier.AddMultipleMatchesInTournament(ctx, tournamentID, matchIDs); err != nil {
	// 	return customerrors.HandleErrMsg(err, "tournament", "error when adding matches in tournament")
	// }

	// Add matches in group
	if err := s.tournamentQuerier.AddMultipleMatchesInTournamentGroup(ctx, groupID, tournamentID, matchIDs); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when adding matches in group")
	}

	return nil
}
