package utils

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/models"
)

func IntermediateAvailability(usersAvailabilitySliceDTO [][]*models.GetDailyAvailabilityByIDDTORes) []*models.GetDailyAvailabilityByIDDTORes {
	if len(usersAvailabilitySliceDTO) != 2 {
		return nil // o manejar el error apropiadamente
	}

	daysOrder := []models.DAY{
		models.DAY_SUNDAY,
		models.DAY_MONDAY,
		models.DAY_TUESDAY,
		models.DAY_WEDNESDAY,
		models.DAY_THURSDAY,
		models.DAY_FRIDAY,
		models.DAY_SATURDAY,
	}

	// Crear mapas para acceder rápidamente a la disponibilidad por día
	availabilityOneMap := make(map[models.DAY]*models.GetDailyAvailabilityByIDDTORes)
	availabilityTwoMap := make(map[models.DAY]*models.GetDailyAvailabilityByIDDTORes)

	for _, daily := range usersAvailabilitySliceDTO[0] {
		availabilityOneMap[daily.Day] = daily
	}
	for _, daily := range usersAvailabilitySliceDTO[1] {
		availabilityTwoMap[daily.Day] = daily
	}

	intermediateAvailability := make([]*models.GetDailyAvailabilityByIDDTORes, 7)

	for i, day := range daysOrder {
		dailyOne, existsOne := availabilityOneMap[day]
		dailyTwo, existsTwo := availabilityTwoMap[day]

		if !existsOne || !existsTwo {
			// Manejar el caso en que falta información para un día
			intermediateAvailability[i] = &models.GetDailyAvailabilityByIDDTORes{
				Day:       day,
				TimeSlots: []*models.GetDailyTimeSlotByIDDTORes{}, // o algún valor predeterminado
			}
			continue
		}

		intermediateAvailability[i] = &models.GetDailyAvailabilityByIDDTORes{
			Day:       day,
			TimeSlots: IntermediateTimeSlots(dailyOne.TimeSlots, dailyTwo.TimeSlots),
		}
	}

	return intermediateAvailability
}

func IntermediateTimeSlots(timeSlotsOne, timeSlotsTwo []*models.GetDailyTimeSlotByIDDTORes) []*models.GetDailyTimeSlotByIDDTORes {
	intermediateTimeSlots := make([]*models.GetDailyTimeSlotByIDDTORes, len(timeSlotsOne))
	for i, slotOne := range timeSlotsOne {
		intermediateTimeSlots[i] = &models.GetDailyTimeSlotByIDDTORes{
			TimeSlot: slotOne.TimeSlot,
			Status:   CombineAvailabilityStatuses(slotOne.Status, timeSlotsTwo[i].Status),
		}
	}

	return intermediateTimeSlots
}

func CombineAvailabilityStatuses(status1, status2 models.AVAILABILITY_STATUS) models.AVAILABILITY_STATUS {
	// Si alguno de los estados es NOT_AVAILABLE, el resultado es NOT_AVAILABLE
	if status1 == models.AVAILABILITY_STATUS_NOT_AVAILABLE || status2 == models.AVAILABILITY_STATUS_NOT_AVAILABLE {
		return models.AVAILABILITY_STATUS_NOT_AVAILABLE
	}

	// Si ambos estados son AVAILABLE, el resultado es AVAILABLE
	if status1 == models.AVAILABILITY_STATUS_AVAILABLE && status2 == models.AVAILABILITY_STATUS_AVAILABLE {
		return models.AVAILABILITY_STATUS_AVAILABLE
	}

	// En cualquier otro caso (que incluye todas las combinaciones con POSSIBLY_AVAILABLE),
	// el resultado es POSSIBLY_AVAILABLE
	return models.AVAILABILITY_STATUS_POSSIBLY_AVAILABLE
}

func OrderAvailability(availabilitySlice []*models.GetDailyAvailabilityByIDDTORes) []*models.GetDailyAvailabilityByIDDTORes {
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
	availabilityMap := make(map[models.DAY]*models.GetDailyAvailabilityByIDDTORes)
	for _, daily := range availabilitySlice {
		availabilityMap[daily.Day] = daily
	}

	orderedAvailability := make([]*models.GetDailyAvailabilityByIDDTORes, 7)

	for i, day := range daysOrder {
		if daily, exists := availabilityMap[day]; exists {
			orderedAvailability[i] = daily
		} else {
			// Si no existe la disponibilidad para este día, crear una entrada vacía
			orderedAvailability[i] = &models.GetDailyAvailabilityByIDDTORes{
				Day:       day,
				TimeSlots: []*models.GetDailyTimeSlotByIDDTORes{},
			}
		}
	}

	return orderedAvailability
}
