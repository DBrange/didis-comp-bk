package utils

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

func IntermediateAvailability(usersAvailabilitySliceDTO [][]*dto.GetDailyAvailabilityByIDDTORes) []*dto.GetDailyAvailabilityByIDDTORes {
	if len(usersAvailabilitySliceDTO) == 1 {
		// Si solo hay un conjunto de disponibilidad, devolverlo directamente
		return usersAvailabilitySliceDTO[0]
	}

	if len(usersAvailabilitySliceDTO) != 2 {
		return nil // o manejar el error apropiadamente
	}

	// Mapea la disponibilidad por día
	availabilityMaps := make([]map[models.DAY]*dto.GetDailyAvailabilityByIDDTORes, len(usersAvailabilitySliceDTO))
	for i, availabilitySlice := range usersAvailabilitySliceDTO {
		availabilityMaps[i] = mapAvailabilityByDay(availabilitySlice)
	}

	// Verifica si alguna de las partes tiene todos los slots como NOT_AVAILABLE
	allNotAvailable := make([]bool, len(usersAvailabilitySliceDTO))
	for i, availabilitySlice := range usersAvailabilitySliceDTO {
		allNotAvailable[i] = allTimeSlotsNotAvailable(availabilitySlice)
	}

	if allNotAvailable[0] && !allNotAvailable[1] {
		return usersAvailabilitySliceDTO[1]
	}
	if allNotAvailable[1] && !allNotAvailable[0] {
		return usersAvailabilitySliceDTO[0]
	}
	if allNotAvailable[0] && allNotAvailable[1] {
		return nil
	}

	// Combina las disponibilidades
	return combineAvailabilities(availabilityMaps[0], availabilityMaps[1])
}

// Mapea la disponibilidad por día
func mapAvailabilityByDay(dailyAvailabilities []*dto.GetDailyAvailabilityByIDDTORes) map[models.DAY]*dto.GetDailyAvailabilityByIDDTORes {
	availabilityMap := make(map[models.DAY]*dto.GetDailyAvailabilityByIDDTORes)
	for _, daily := range dailyAvailabilities {
		availabilityMap[daily.Day] = daily
	}
	return availabilityMap
}

// Verifica si todos los slots en la disponibilidad son NOT_AVAILABLE
func allTimeSlotsNotAvailable(dailyAvailabilities []*dto.GetDailyAvailabilityByIDDTORes) bool {
	for _, daily := range dailyAvailabilities {
		for _, timeSlot := range daily.TimeSlots {
			if timeSlot.Status != models.AVAILABILITY_STATUS_NOT_AVAILABLE {
				return false
			}
		}
	}
	return true
}

// Combina dos disponibilidades en una
func combineAvailabilities(availabilityOneMap, availabilityTwoMap map[models.DAY]*dto.GetDailyAvailabilityByIDDTORes) []*dto.GetDailyAvailabilityByIDDTORes {
	daysOrder := []models.DAY{
		models.DAY_SUNDAY,
		models.DAY_MONDAY,
		models.DAY_TUESDAY,
		models.DAY_WEDNESDAY,
		models.DAY_THURSDAY,
		models.DAY_FRIDAY,
		models.DAY_SATURDAY,
	}

	intermediateAvailability := make([]*dto.GetDailyAvailabilityByIDDTORes, len(daysOrder))
	allNotAvailable := true // Bandera para verificar si todos los slots son NOT_AVAILABLE

	for i, day := range daysOrder {
		dailyOne, existsOne := availabilityOneMap[day]
		dailyTwo, existsTwo := availabilityTwoMap[day]

		if !existsOne || !existsTwo {
			intermediateAvailability[i] = &dto.GetDailyAvailabilityByIDDTORes{
				Day:       day,
				TimeSlots: []*dto.GetDailyTimeSlotByIDDTORes{}, // o algún valor predeterminado
			}
			continue
		}

		combinedTimeSlots := IntermediateTimeSlots(dailyOne.TimeSlots, dailyTwo.TimeSlots)
		intermediateAvailability[i] = &dto.GetDailyAvailabilityByIDDTORes{
			Day:       day,
			TimeSlots: combinedTimeSlots,
		}

		// Verificar si alguno de los slots no es NOT_AVAILABLE
		for _, slot := range combinedTimeSlots {
			if slot.Status != models.AVAILABILITY_STATUS_NOT_AVAILABLE {
				allNotAvailable = false
			}
		}
	}

	// Si todos los slots son NOT_AVAILABLE, devolver nil
	if allNotAvailable {
		return nil
	}

	return intermediateAvailability
}


// Combina los estados de disponibilidad
func IntermediateTimeSlots(timeSlotsOne, timeSlotsTwo []*dto.GetDailyTimeSlotByIDDTORes) []*dto.GetDailyTimeSlotByIDDTORes {
	intermediateTimeSlots := make([]*dto.GetDailyTimeSlotByIDDTORes, len(timeSlotsOne))
	for i, slotOne := range timeSlotsOne {
		intermediateTimeSlots[i] = &dto.GetDailyTimeSlotByIDDTORes{
			TimeSlot: slotOne.TimeSlot,
			Status:   combineAvailabilityStatuses(slotOne.Status, timeSlotsTwo[i].Status),
		}
	}
	return intermediateTimeSlots
}

// Combina los estados de disponibilidad
func combineAvailabilityStatuses(status1, status2 models.AVAILABILITY_STATUS) models.AVAILABILITY_STATUS {
	if status1 == models.AVAILABILITY_STATUS_NOT_AVAILABLE || status2 == models.AVAILABILITY_STATUS_NOT_AVAILABLE {
		return models.AVAILABILITY_STATUS_NOT_AVAILABLE
	}

	if status1 == models.AVAILABILITY_STATUS_AVAILABLE && status2 == models.AVAILABILITY_STATUS_AVAILABLE {
		return models.AVAILABILITY_STATUS_AVAILABLE
	}

	return models.AVAILABILITY_STATUS_POSSIBLY_AVAILABLE
}

// Ordena por si vienen desordenados
func OrderAvailability(availabilitySlice []*dto.GetDailyAvailabilityByIDDTORes) []*dto.GetDailyAvailabilityByIDDTORes {
	daysOrder := []models.DAY{
		models.DAY_SUNDAY,
		models.DAY_MONDAY,
		models.DAY_TUESDAY,
		models.DAY_WEDNESDAY,
		models.DAY_THURSDAY,
		models.DAY_FRIDAY,
		models.DAY_SATURDAY,
	}

	// Crear un mapa para acceder rápidamente a la disponibilidad por día
	availabilityMap := make(map[models.DAY]*dto.GetDailyAvailabilityByIDDTORes)
	for _, daily := range availabilitySlice {
		availabilityMap[daily.Day] = daily
	}

	orderedAvailability := make([]*dto.GetDailyAvailabilityByIDDTORes, 7)

	for i, day := range daysOrder {
		if daily, exists := availabilityMap[day]; exists {
			orderedAvailability[i] = daily
		} else {
			// Si no existe la disponibilidad para este día, crear una entrada vacía
			orderedAvailability[i] = &dto.GetDailyAvailabilityByIDDTORes{
				Day:       day,
				TimeSlots: []*dto.GetDailyTimeSlotByIDDTORes{},
			}
		}
	}

	return orderedAvailability
}
