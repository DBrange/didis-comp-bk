package services

import (
	"context"

	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
)

func (s *TournamentService) OrganizeBracket(ctx context.Context,tournamentID string, competitorsDTOs []*dto.UpdateCompetitorMatchDTOReq) error {
	competitorIDs := make([]string, len(competitorsDTOs))
	matchesIDs := make([]string, 0, len(competitorsDTOs)/2)

	// Divide los partidos de los competidores para luego verificar si existen
	for i, competitorsDTO := range competitorsDTOs {
		competitorIDs[i] = *competitorsDTO.CompetitorID

		if !utils.ContainsID(matchesIDs, *competitorsDTO.MatchID) {
			matchesIDs = append(matchesIDs, *competitorsDTO.MatchID)
		}
	}

	for _, competitorsDTO := range competitorsDTOs {
		if err := s.tournamentQueryer.AddMatchInCompetitorStats(ctx, *competitorsDTO.CompetitorID, *competitorsDTO.MatchID); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when updating competitor stats")
		}
	}

	if err := s.tournamentQueryer.VerifyMatchesExist(ctx, matchesIDs); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when verifying matches")
	}

	// if err := s.tournamentQueryer.VerifyMultipleCompetitorsExists(ctx, competitorIDs); err != nil {
	// 	return customerrors.HandleErrMsg(err, "tournament", "error when verifying competitors")
	// }

	if err := s.tournamentQueryer.VerifyMultipleCompetitorsExistsInTournament(ctx, tournamentID, competitorIDs); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when verifying competitors")
	}

	if err := s.tournamentQueryer.UpdateMultipleCompetitorMatches(ctx, competitorsDTOs); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when updating competitorMatches")
	}

	return nil
}

// TODA ESTA LOGICA LA VOY A NECESITAR EN EL FRONTEND !!!

// ESTA FUNCION PUEDE QUE ME SIRVA PARA EL FRONT, SI NECESITO LOS DATOS QUE USABA ACA ---> s.tournamentQueryer.GetPositionsBracketMatch

// func (s *TournamentService) OrganizeBracket(ctx context.Context, roundID string, organizeType models.ORGANIZE_TYPE, organizeBracket models.ORGANIZE_BRACKET) error {
// 	positionsDTO, err := s.tournamentQueryer.GetPositionsBracketMatch(ctx, roundID, organizeType, organizeBracket)
// 	if err != nil {
// 		return customerrors.HandleErrMsg(err, "tournament", "error when getting postions in matches and competitors")
// 	}

// 	if organizeType == models.ORGANIZE_TYPE_COMPETITOR && organizeBracket == models.ORGANIZE_BRACKET_RANK {
// 		s.sortByRank(positionsDTO, organizeType, organizeBracket)
// 	}

// 	return nil
// }

// func (s *TournamentService) sortByRank(positionsDTO []*dto.GetPositionsBracketMatchDTORes, organizeType models.ORGANIZE_TYPE, organizeBracket models.ORGANIZE_BRACKET) (any, error) {
// 	positionsSoted := make([]*dto.GetPositionsBracketMatchDTORes, len(positionsDTO))
// 	competitors := make([]*dto.GetPositionsBracketMatchCompetitorDTORes, len(positionsDTO)*2)

// 	for _, positionDTO := range positionsDTO {
// 		competitors = append(competitors, positionDTO.Competitors...)
// 	}

// 	competitorsSorted, nilCompetitorsSorted := sortCompetitors(competitors)
// 	matchesSorted := sortMatches(positionsDTO)

// 	nose(matchesSorted, competitorsSorted, nilCompetitorsSorted)

// 	return positionsSoted, nil
// }

// func sortCompetitors(competitors []*dto.GetPositionsBracketMatchCompetitorDTORes) ([]*dto.GetPositionsBracketMatchCompetitorDTORes, []*dto.GetPositionsBracketMatchCompetitorDTORes) {
// 	// Separar competidores nil y no nil
// 	var validCompetitors, nilCompetitors []*dto.GetPositionsBracketMatchCompetitorDTORes

// 	for _, competitor := range competitors {
// 		if competitor.CurrentPosition != nil {
// 			validCompetitors = append(validCompetitors, competitor)
// 		} else {
// 			nilCompetitors = append(nilCompetitors, competitor)
// 		}
// 	}

// 	// Ordenar competidores válidos
// 	sort.Slice(validCompetitors, func(i, j int) bool {
// 		return *validCompetitors[i].CurrentPosition < *validCompetitors[j].CurrentPosition
// 	})

// 	return validCompetitors, nilCompetitors
// }

// func sortMatches(matches []*dto.GetPositionsBracketMatchDTORes) []*dto.GetPositionsBracketMatchDTORes {
// 	matchedSorted := make([]*dto.GetPositionsBracketMatchDTORes, len(matches))

// 	for i, match := range matches {
// 		matchedSorted[i] = &dto.GetPositionsBracketMatchDTORes{
// 			ID:            match.ID,
// 			PositionMatch: match.PositionMatch,
// 		}
// 	}

// 	// Ordenar el array clonado
// 	sort.Slice(matchedSorted, func(i, j int) bool {
// 		return matchedSorted[i].PositionMatch < matchedSorted[j].PositionMatch
// 	})

// 	return matchedSorted
// }

// func nose(matches []*dto.GetPositionsBracketMatchDTORes, competitors []*dto.GetPositionsBracketMatchCompetitorDTORes, nilCompetitors []*dto.GetPositionsBracketMatchCompetitorDTORes) {
// 	// Top seeds
// 	firstBestComp := competitors[0]
// 	secondBestComp := competitors[1]

// 	firstSeed := matches[0].Competitors[0]
// 	secondSeed := matches[len(competitors)-1].Competitors[1]

// 	// añadiendo a al primer serbrado
// 	firstSeed.Position = 1
// 	firstSeed.ID = firstBestComp.ID

// 	// añadiendo a al segundo serbrado
// 	secondSeed.Position = 2
// 	secondSeed.ID = secondBestComp.ID

// 	addTopSeeds(matches, competitors)

// }

// func addTopSeeds(matches []*dto.GetPositionsBracketMatchDTORes, competitors []*dto.GetPositionsBracketMatchCompetitorDTORes) []*dto.GetPositionsBracketMatchDTORes {
// 	if len(matches) <= 2 {
// 		return matches
// 	}

// 	queue := []int{0, len(matches)} // Almacena índices en lugar de slices
// 	currentComp := 2

// 	for len(queue) > 0 && currentComp < len(competitors) {
// 		// Tomar el siguiente segmento de la cola
// 		start := queue[0]
// 		end := queue[1]
// 		queue = queue[2:]

// 		if end-start <= 2 {
// 			continue
// 		}

// 		// Encontrar el punto medio
// 		mid := (start + end) / 2

// 		// Asignar los competidores de mejor posición a los seeds
// 		firstSeed := matches[mid-1].Competitors[1]
// 		secondSeed := matches[mid].Competitors[0]

// 		firstBestComp := competitors[currentComp]
// 		secondBestComp := competitors[currentComp+1]

// 		firstSeed.Position = 2
// 		firstSeed.ID = firstBestComp.ID

// 		secondSeed.Position = 1
// 		secondSeed.ID = secondBestComp.ID

// 		currentComp += 2

// 		// Agregar las nuevas mitades a la cola para procesar
// 		// se agregan 4, ya que luego con el slice se eliminan las primeras 2 y quedan las 2 sioguientes, y asi.
// 		queue = append(queue, start, mid, mid, end)
// 	}

// 	return matches
// }
