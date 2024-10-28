package services

import (
	"context"
	"fmt"
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	api_util "github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/cmd/utils"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	customerrors "github.com/DBrange/didis-comp-bk/pkg/custom_errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *TournamentService) OrganizeBracket(ctx context.Context, tournamentID string, competitorMatchDTOs []*dto.UpdateCompetitorMatchDTOReq) error {
	// 	categoryID, err := s.tournamentQuerier.GetCategoryIDOfTournament(ctx, tournamentID)
	// if err != nil {
	// 	return customerrors.HandleErrMsg(err, "tournament", "error when getting tournament competitors")
	// }

	
	// competitorsDTO, err := s.tournamentQuerier.GetCompetitorsInTournament(ctx, tournamentID, categoryID, "", 0, true)
	// 	if err != nil {
	// 		return customerrors.HandleErrMsg(err, "tournament", "error when getting tournament competitors")
	// 	}

	
	
	
	competitorIDs := make([]*string, len(competitorMatchDTOs))
	matchesIDs := make([]string, 0, len(competitorMatchDTOs)/2)
	competitorsInMatchMap := make(map[string][]string)

	for i, competitorsDTO := range competitorMatchDTOs {
		// Divide los partidos y los competidores para luego verificar si existen
		competitorIDs[i] = competitorsDTO.CompetitorID

		if !api_util.ContainsID(matchesIDs, competitorsDTO.MatchID) {
			matchesIDs = append(matchesIDs, competitorsDTO.MatchID)
		}

		// Unir los competidores con el mismo match
		if competitorsDTO.CompetitorID != nil {
			competitorsInMatchMap[competitorsDTO.MatchID] = append(competitorsInMatchMap[competitorsDTO.MatchID], *competitorsDTO.CompetitorID)
		}

	}

	err := s.tournamentQuerier.WithTransaction(ctx, func(sessCtx mongo.SessionContext) error {
		// Get tournament matches
		matchesToRemove, err := s.tournamentQuerier.GetTournamentMatchesByID(sessCtx, tournamentID)
		if err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when adding competitors in groups")
		}
		// get all the competitors that are in a match of the tournament and then remove the match from their statistics,
		// since the same matches will be created again
		competitorsForRemoveMatches, err := s.tournamentQuerier.GetCompetitorIDsFromMatches(sessCtx, matchesToRemove)
		if err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when adding competitors in groups")
		}

		if err := s.eliminateBracketMatches(sessCtx, competitorsForRemoveMatches, matchesToRemove); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when eliminate group")
		}


		for _, competitorsDTO := range competitorMatchDTOs {
			if competitorsDTO.CompetitorID != nil {
				if err := s.tournamentQuerier.AddMatchInCompetitorStats(sessCtx, *competitorsDTO.CompetitorID, competitorsDTO.MatchID); err != nil {
					return customerrors.HandleErrMsg(err, "tournament", "error when updating competitor stats")
				}
			}
		}

		if err := s.tournamentQuerier.VerifyMatchesExist(sessCtx, matchesIDs); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when verifying matches")
		}
		var competitorIDsStr []string
		for _, id := range competitorIDs {
			if id != nil {
				competitorIDsStr = append(competitorIDsStr, *id) // Desreferenciar el puntero
			}
		}
		if err := s.tournamentQuerier.VerifyMultipleCompetitorsExistsInTournament(sessCtx, tournamentID, competitorIDsStr); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when verifying competitors")
		}

		if err := s.tournamentQuerier.UpdateMultipleCompetitorMatches(sessCtx, competitorMatchDTOs); err != nil {
			return customerrors.HandleErrMsg(err, "tournament", "error when updating competitorMatches")
		}

		courtAvailability, tournamentAvailabilities, err := s.getPartialAvailabilityInTournament(sessCtx, tournamentID)
		if err != nil {
			return err
		}

		timetablesNotAvailables := []time.Time{}
		if err := s.updateMatchesDates(sessCtx, competitorsInMatchMap, courtAvailability, tournamentAvailabilities, timetablesNotAvailables); err != nil {
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
	tournamentAvailabilities []*models.GetDailyAvailabilityByIDDTORes,
	timetablesNotAvailables *[]time.Time,
	availableCourts int,
	hoursForMatch time.Duration,
) ([]*dto.MatchDateDTOReq, error) {
	matchesDate := make([]*dto.MatchDateDTOReq, 0, len(competitorsInMatchMap))
for _, v := range competitorsInMatchMap{
	fmt.Printf("el tamaño: %+v",v)

}
	for matchID, competitorIDs := range competitorsInMatchMap {
		if len(competitorIDs) != 2 {
			matchDate := &dto.MatchDateDTOReq{
				ID:   matchID,
				Date: nil, // Establecer como nil si matchTime es time.Time{}
			}
			matchesDate = append(matchesDate, matchDate)
		}

		if len(competitorIDs) != 2 {
			continue
		}

		availabilityDTOs, err := s.tournamentQuerier.GetMultipleAvailabilitiesByCompetitor(ctx, competitorIDs)
		if err != nil {
			return nil, err
		}

		// Crear disponibilidad intermedia de competidores y canchas
		intermediateAvailabilityCompetitorDTO := utils.IntermediateAvailability(availabilityDTOs)
		iacSlice := [][]*models.GetDailyAvailabilityByIDDTORes{intermediateAvailabilityCompetitorDTO, tournamentAvailabilities}
		intermediateAvailabilityCourtDTO := utils.IntermediateAvailability(iacSlice)


		// Buscar el primer `TimeSlot` AVAILABLE o POSSIBLY_AVAILABLE
		matchTime, err := findFirstAvailableTimeSlot(intermediateAvailabilityCourtDTO, *timetablesNotAvailables, availableCourts)
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
		*timetablesNotAvailables = append(*timetablesNotAvailables, matchTime, matchTime.Add(hoursForMatch*time.Hour))
	}

	return matchesDate, nil
}

func (s *TournamentService) getPartialAvailabilityInTournament(ctx context.Context, tournamentID string) (*dto.TournamentAvailabilityDTO, []*models.GetDailyAvailabilityByIDDTORes, error) {
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

func (s *TournamentService) updateMatchesDates(ctx context.Context, competitorsInMatchMap map[string][]string, courtAvailability *dto.TournamentAvailabilityDTO, tournamentAvailabilities []*models.GetDailyAvailabilityByIDDTORes, timetablesNotAvailables []time.Time) error {
	hoursForMatch := time.Duration(courtAvailability.AverageHours - 1)
	matchesDate, err := s.ScheduleMatches(ctx, competitorsInMatchMap, tournamentAvailabilities, &timetablesNotAvailables, courtAvailability.AvailableCourts, hoursForMatch)
	if err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when set date in matches")
	}


	
	// actualizar todos los matches con su nuevo date
	if err := s.tournamentQuerier.UpdateMultipleMatchesDate(ctx, matchesDate); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when updating match date")
	}

	return nil
}

func findFirstAvailableTimeSlot(availability []*models.GetDailyAvailabilityByIDDTORes, timetablesNotAvailables []time.Time, availableCourts int) (time.Time, error) {
	if availability == nil {
		return time.Time{}, nil
	}

	now := time.Now()
	startDate := now.Add(24 * time.Hour) // Comienza a buscar a partir de mañana
	maxDays := 16                        // Límite de días para buscar a partir de startDate

	for dayOffset := 0; dayOffset < maxDays; dayOffset++ { // Evalúa hasta un máximo de 16 días
		searchDate := startDate.AddDate(0, 0, dayOffset)

		for i := 0; i < 2; i++ { // Primero busca AVAILABLE, luego POSSIBLY_AVAILABLE
			desiredStatus := models.AVAILABILITY_STATUS_AVAILABLE
			if i == 1 {
				desiredStatus = models.AVAILABILITY_STATUS_POSSIBLY_AVAILABLE
			}

			for _, dailyAvailability := range availability {
				dayOfWeek := convertDayToString(dailyAvailability.Day)
				weekday := parseWeekday(dayOfWeek)

				// Verifica si el `searchDate` coincide con el día de la semana deseado
				if searchDate.Weekday() == weekday {
					for _, timeSlot := range dailyAvailability.TimeSlots {
						if timeSlot.Status == desiredStatus {
							slotTime, _ := time.Parse("15:04", timeSlot.TimeSlot)
							fullSlotTime := time.Date(searchDate.Year(), searchDate.Month(), searchDate.Day(), slotTime.Hour(), slotTime.Minute(), 0, 0, startDate.Location())
							
							fmt.Printf("Fecha evaluada: %+v\n", fullSlotTime)
							
							// Verifica que no esté en la lista de `timetablesNotAvailables`
							if isTimeSlotAvailable(fullSlotTime, timetablesNotAvailables, availableCourts) {
								return fullSlotTime, nil
							}
						}
					}
				}
			}
		}
	}

	return time.Time{}, nil
}






// Convierte el nombre del día a un valor de `time.Weekday`
func parseWeekday(day string) time.Weekday {
	switch day {
	case "Sunday":
		return time.Sunday
	case "Monday":
		return time.Monday
	case "Tuesday":
		return time.Tuesday
	case "Wednesday":
		return time.Wednesday
	case "Thursday":
		return time.Thursday
	case "Friday":
		return time.Friday
	case "Saturday":
		return time.Saturday
	default:
		return time.Sunday // Valor predeterminado
	}
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

func (s *TournamentService) eliminateBracketMatches(ctx context.Context,  competitorIDs, matchesToRemove []string) error {
	// sacar los matches de competitorStats y cambiar las
	if err := s.tournamentQuerier.RemoveMultipleCompetitorStatsMatches(ctx, competitorIDs, matchesToRemove); err != nil {
		return customerrors.HandleErrMsg(err, "tournament", "error when adding competitors in groups")
	}

	return nil
}