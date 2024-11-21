package mappers

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/avaliability/dao"
)

// Función para mapear GetDayTimeSlotDAORes a GetDayTimeSlotDTORes
func GetUsersAvailabilityDAOtoDTO(availabilityOIDs []*dao.GetDayTimeSlotDAORes) []*dto.GetDayTimeSlotDTORes {
	var result []*dto.GetDayTimeSlotDTORes

	for _, availability := range availabilityOIDs {
		dtoRes := mapDayTimeSlotDAOtoDTO(availability)
		result = append(result, dtoRes)
	}

	return result
}

// Función para mapear un GetDayTimeSlotDAORes a GetDayTimeSlotDTORes
func mapDayTimeSlotDAOtoDTO(availability *dao.GetDayTimeSlotDAORes) *dto.GetDayTimeSlotDTORes {
	return &dto.GetDayTimeSlotDTORes{
		ID:                  availability.ID.Hex(),
		DailyAvailabilities: mapDailyAvailabilitiesDAOtoDTO(availability.DailyAvailabilities),
	}
}

// Función para mapear los DailyAvailabilities
func mapDailyAvailabilitiesDAOtoDTO(dailyAvailabilities []*dao.GetDailyAvailabilityByIDDAORes) []*models.GetDailyAvailabilityByIDDTORes {
	var result []*models.GetDailyAvailabilityByIDDTORes

	for _, daily := range dailyAvailabilities {
		result = append(result, &models.GetDailyAvailabilityByIDDTORes{
			Day:       daily.Day,
			TimeSlots: mapTimeSlotsDAOtoDTO(daily.TimeSlots),
		})
	}

	return result
}

// Función para mapear los TimeSlots
func mapTimeSlotsDAOtoDTO(timeSlots []*dao.GetDailyTimeSlotByIDDAORes) []*models.GetDailyTimeSlotByIDDTORes {
	var result []*models.GetDailyTimeSlotByIDDTORes

	for _, slot := range timeSlots {
		result = append(result, &models.GetDailyTimeSlotByIDDTORes{
			TimeSlot: slot.TimeSlot,
			Status:   slot.Status,
		})
	}

	return result
}
