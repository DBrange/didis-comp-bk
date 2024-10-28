package mappers

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/avaliability/dao"
)

func GetAvailabilityDailySliceDTOtoDAO(dailyAvailabilitiesDAO []*dao.GetDailyAvailabilityByIDDAORes) []*models.GetDailyAvailabilityByIDDTORes {
	dailyAvailabilitiesDTO := make([]*models.GetDailyAvailabilityByIDDTORes, len(dailyAvailabilitiesDAO))

	for i, daily := range dailyAvailabilitiesDAO {
		dailyAvailabilitiesDTO[i] = &models.GetDailyAvailabilityByIDDTORes{
			Day:       daily.Day,
			TimeSlots: GetAvailabilityDailyTimeSlotSliceDTOtoDAO(daily.TimeSlots),
		}
	}

	return dailyAvailabilitiesDTO
}

func GetAvailabilityDailyTimeSlotSliceDTOtoDAO(timeSlots []*dao.GetDailyTimeSlotByIDDAORes) []*models.GetDailyTimeSlotByIDDTORes {
	timeSlotsDTO := make([]*models.GetDailyTimeSlotByIDDTORes, len(timeSlots))

	for i, slot := range timeSlots {
		timeSlotsDTO[i] = &models.GetDailyTimeSlotByIDDTORes{
			TimeSlot: slot.TimeSlot,
			Status:   slot.Status,
		}
	}

	return timeSlotsDTO
}
