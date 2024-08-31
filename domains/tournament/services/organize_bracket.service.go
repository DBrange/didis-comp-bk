package services

import (
	"context"
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	tournament_utils "github.com/DBrange/didis-comp-bk/domains/tournament/utils"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *TournamentService) OrganizeBracket(ctx context.Context, tournamentID string, competitorsDTOs []*dto.UpdateCompetitorMatchDTOReq) error {
	competitorIDs := make([]string, len(competitorsDTOs))
	matchesIDs := make([]string, 0, len(competitorsDTOs)/2)
	competitorsInMatchMap := make(map[string][]string)

	for i, competitorsDTO := range competitorsDTOs {
		// Divide los partidos y los competidores para luego verificar si existen
		competitorIDs[i] = competitorsDTO.CompetitorID

		if !utils.ContainsID(matchesIDs, competitorsDTO.MatchID) {
			matchesIDs = append(matchesIDs, competitorsDTO.MatchID)
		}

		// Unir los competidores con el mismo match
		competitorsInMatchMap[competitorsDTO.MatchID] = append(competitorsInMatchMap[competitorsDTO.MatchID], competitorsDTO.CompetitorID)

	}

	err := s.tournamentQuerier.WithTransaction(ctx, func(sessCtx mongo.SessionContext) error {
		for _, competitorsDTO := range competitorsDTOs {
			if err := s.tournamentQuerier.AddMatchInCompetitorStats(ctx, competitorsDTO.CompetitorID, competitorsDTO.MatchID); err != nil {
				return customerrors.HandleErrMsg(err, "tournament", "error when updating competitor stats")
			}
		}

		if err := s.tournamentQuerier.VerifyMatchesExist(ctx, matchesIDs); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when verifying matches")
		}

		if err := s.tournamentQuerier.VerifyMultipleCompetitorsExistsInTournament(ctx, tournamentID, competitorIDs); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when verifying competitors")
		}

		if err := s.tournamentQuerier.UpdateMultipleCompetitorMatches(ctx, competitorsDTOs); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when updating competitorMatches")
		}

		courtAvailability, tournamentAvailabilities, err := s.getPartialAvailabilityInTournament(ctx, tournamentID)
		if err != nil {
			return err
		}

		timetablesNotAvailables := []time.Time{}
		if err := s.updateMatchesDates(ctx, competitorsInMatchMap, courtAvailability, tournamentAvailabilities, timetablesNotAvailables); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when updating competitorMatches")
		}

		return nil
	})
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when creating tournament location")
	}

	return nil
}

// crear funcionalidad en donde tomando como partida el dia de hoy, se seleccione el timeslot que diga AVAILABLE mas cercano, pero a partir de mañana, no del mismo dia
// guardar el matchID y el timeslot que es ta representado como "00:00" como un time.Time, asi se guarda como un Date.
// luego depositar ese timeslot como time.Time dentro de timetablesNotAvailables
// ese timeslot siempre debe representar como si fueran 2 horarios ocupados, si por ejemplo es 00:00, debe ocupar el timeslot de 00:00 y 01:00, o sea que siempre en el slice de timetablesNotAvailables se debe guardar el timeslot que corresponde, y el que le sigue
// luego en el siguiente recorrido se debera verificar si ese timeslot ya esta ocupado la cantidad de veces que dice courts.AvailableCourts
// en caso de estar ocupado, se debe buscar el siguiente timeslot que este como AVAILABLE
// en caso de que al recorrer toda la semana no haya ningun AVAILABLE disponible se debera volver a recorrer todos los timeslot pero en busca de POSSIBLE_AVAILABLE, con toda la misma logica de antes
// en caso que tampoco haya coincidencia, se debera volver a la logica de buscar AVAILABLE pero de la siguiente semana, y seguir con el mismo patron de antes
func (s *TournamentService) ScheduleMatches(
	ctx context.Context,
	competitorsInMatchMap map[string][]string,
	tournamentAvailabilities []*dto.GetDailyAvailabilityByIDDTORes,
	timetablesNotAvailables []time.Time,
	availableCourts int,
	hoursForMatch time.Duration,
) ([]*dto.MatchDateDTOReq, error) {
	matchesDate := make([]*dto.MatchDateDTOReq, 0, len(competitorsInMatchMap))

	for matchID, competitorIDs := range competitorsInMatchMap {
		availabilityDTOs, err := s.tournamentQuerier.GetMultipleAvailabilitiesByCompetitor(ctx, competitorIDs)
		if err != nil {
			return nil, err
		}

		// Crear disponibilidad intermedia de competidores y canchas
		intermediateAvailabilityCompetitorDTO := tournament_utils.IntermediateAvailability(availabilityDTOs)

		iacSlice := [][]*dto.GetDailyAvailabilityByIDDTORes{intermediateAvailabilityCompetitorDTO, tournamentAvailabilities}
		intermediateAvailabilityCourtDTO := tournament_utils.IntermediateAvailability(iacSlice)

		// Buscar el primer `TimeSlot` AVAILABLE o POSSIBLY_AVAILABLE
		matchTime, err := findFirstAvailableTimeSlot(intermediateAvailabilityCourtDTO, timetablesNotAvailables, availableCourts)
		if err != nil {
			return nil, err
		}

		// Guardar el `TimeSlot` encontrado
		var matchDate *dto.MatchDateDTOReq
		if matchTime.IsZero() {
			matchDate = &dto.MatchDateDTOReq{
				ID:   matchID,
				Date: nil, // Establecer como nil si matchTime es time.Time{}
			}
		} else {
			matchDate = &dto.MatchDateDTOReq{
				ID:   matchID,
				Date: &matchTime, // Asignar la dirección de matchTime
			}
		}
		matchesDate = append(matchesDate, matchDate)

		// Agregar el `TimeSlot` y el siguiente a `timetablesNotAvailables`
		timetablesNotAvailables = append(timetablesNotAvailables, matchTime, matchTime.Add(hoursForMatch*time.Hour))
	}

	return matchesDate, nil
}

func (s *TournamentService) getPartialAvailabilityInTournament(ctx context.Context, tournamentID string) (*dto.TournamentAvailabilityDTO, []*dto.GetDailyAvailabilityByIDDTORes, error) {
	courtAvailability, err := s.tournamentQuerier.GetTournamentAvailavility(ctx, tournamentID)
	if err != nil {
		return nil, nil, customerrors.HandleErrMsg(err, "tournament", "error when updating competitorMatches")
	}

	tournamentAvailabilities, err := s.tournamentQuerier.GetAvailabilityByTournamentID(ctx, tournamentID)
	if err != nil {
		return nil, nil, customerrors.HandleErrMsg(err, "tournament", "error when updating competitorMatches")
	}

	return courtAvailability, tournamentAvailabilities, nil
}

func (s *TournamentService) updateMatchesDates(ctx context.Context, competitorsInMatchMap map[string][]string, courtAvailability *dto.TournamentAvailabilityDTO, tournamentAvailabilities []*dto.GetDailyAvailabilityByIDDTORes, timetablesNotAvailables []time.Time) error {
	// se le resta la hora que ya ocuparia por defecto
	hoursForMatch := time.Duration(courtAvailability.AverageHours - 1)
	matchesDate, err := s.ScheduleMatches(ctx, competitorsInMatchMap, tournamentAvailabilities, timetablesNotAvailables, courtAvailability.AvailableCourts, hoursForMatch)
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when set date in matches")
	}

	// actualizar todos los matches con su nuevo date
	if err := s.tournamentQuerier.UpdateMultipleMatchesDate(ctx, matchesDate); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when updating match date")
	}

	return nil
}

func findFirstAvailableTimeSlot(availability []*dto.GetDailyAvailabilityByIDDTORes, timetablesNotAvailables []time.Time, availableCourts int) (time.Time, error) {
	if availability == nil {
		return time.Time{}, nil
	}

	now := time.Now()
	startDate := now.Add(24 * time.Hour) // Comenzar a buscar a partir de mañana

	for i := 0; i < 2; i++ { // Primero buscar AVAILABLE, luego POSSIBLY_AVAILABLE
		desiredStatus := models.AVAILABILITY_STATUS_AVAILABLE
		if i == 1 {
			desiredStatus = models.AVAILABILITY_STATUS_POSSIBLY_AVAILABLE
		}

		for _, dailyAvailability := range availability {
			dayOfWeek := convertDayToString(dailyAvailability.Day)
			parsedDayOfWeek, _ := time.Parse("Monday", dayOfWeek)
			for _, timeSlot := range dailyAvailability.TimeSlots {
				if timeSlot.Status == desiredStatus {
					slotTime, _ := time.Parse("15:04", timeSlot.TimeSlot)
					fullSlotTime := time.Date(startDate.Year(), startDate.Month(), parsedDayOfWeek.Day(), slotTime.Hour(), slotTime.Minute(), 0, 0, startDate.Location())

					// Verificar que no esté en la lista de `timetablesNotAvailables`
					if isTimeSlotAvailable(fullSlotTime, timetablesNotAvailables, availableCourts) {
						return fullSlotTime, nil
					}
				}
			}
		}

		startDate = startDate.AddDate(0, 0, 7) // Moverse a la siguiente semana
	}

	return time.Time{}, nil
}

func isTimeSlotAvailable(timeSlot time.Time, timetablesNotAvailables []time.Time, availableCourts int) bool {
	count := 0
	for _, occupied := range timetablesNotAvailables {
		if occupied == timeSlot {
			count++
			if count >= availableCourts {
				return false
			}
		}
	}
	return true
}

func convertDayToString(day models.DAY) string {
	switch day {
	case models.DAY_SUNDAY:
		return "Sunday"
	case models.DAY_MONDAY:
		return "Monday"
	case models.DAY_TUESDAY:
		return "Tuesday"
	case models.DAY_WEDNESDAY:
		return "Wednesday"
	case models.DAY_THURSDAY:
		return "Thursday"
	case models.DAY_FRIDAY:
		return "Friday"
	case models.DAY_SATURDAY:
		return "Saturday"
	default:
		return ""
	}
}

// TODA ESTA LOGICA LA VOY A NECESITAR EN EL FRONTEND !!!

// ESTA FUNCION PUEDE QUE ME SIRVA PARA EL FRONT, SI NECESITO LOS DATOS QUE USABA ACA ---> s.tournamentQuerier.GetPositionsBracketMatch

// func (s *TournamentService) OrganizeBracket(ctx context.Context, roundID string, organizeType models.ORGANIZE_TYPE, organizeBracket models.ORGANIZE_BRACKET) error {
// 	positionsDTO, err := s.tournamentQuerier.GetPositionsBracketMatch(ctx, roundID, organizeType, organizeBracket)
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
