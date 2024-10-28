package mappers

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/avaliability/dao"
)

func CreateAvailabilityDailySliceDTOtoDAO(dailyAvailabilitiesDTO []*models.GetDailyAvailabilityByIDDTORes) []*dao.CreateDailyAvailability {
	dailyAvailabilitiesDAO := make([]*dao.CreateDailyAvailability, len(dailyAvailabilitiesDTO))

	for i, daily := range dailyAvailabilitiesDTO {
		dailyAvailabilitiesDAO[i] = &dao.CreateDailyAvailability{
			Day:       daily.Day,
			TimeSlots: CreateAvailabilityDailyTimeSlotSliceDTOtoDAO(daily.TimeSlots),
		}
	}

	return dailyAvailabilitiesDAO
}

func CreateAvailabilityDailyTimeSlotSliceDTOtoDAO(timeSlots []*models.GetDailyTimeSlotByIDDTORes) []*dao.CreateTimeSlot {
	timeSlotsDAO := make([]*dao.CreateTimeSlot, len(timeSlots))

	for i, slot := range timeSlots {
		timeSlotsDAO[i] = &dao.CreateTimeSlot{
			TimeSlot: slot.TimeSlot,
			Status:   slot.Status,
		}
	}

	return timeSlotsDAO
}
