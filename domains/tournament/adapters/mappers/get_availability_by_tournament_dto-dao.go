package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/avaliability/dao"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

func GetAvailabilityByTournamentIDDAOtoDTO(dailyAvailabilitiesDAO []*dao.GetDailyAvailabilityByIDDAORes) []*dto.GetDailyAvailabilityByIDDTORes {
	dailyAvailabilitiesDTO := make([]*dto.GetDailyAvailabilityByIDDTORes, len(dailyAvailabilitiesDAO))

	for i, daily := range dailyAvailabilitiesDAO {
		dailyAvailabilitiesDTO[i] = &dto.GetDailyAvailabilityByIDDTORes{
			Day:       daily.Day,
			TimeSlots: GetAvailabilityDailyTimeSlotSliceDTOtoDAO(daily.TimeSlots),
		}
	}

	return dailyAvailabilitiesDTO
}

func GetAvailabilityDailyTimeSlotSliceDTOtoDAO(timeSlots []*dao.GetDailyTimeSlotByIDDAORes) []*dto.GetDailyTimeSlotByIDDTORes {
	timeSlotsDTO := make([]*dto.GetDailyTimeSlotByIDDTORes, len(timeSlots))

	for i, slot := range timeSlots {
		timeSlotsDTO[i] = &dto.GetDailyTimeSlotByIDDTORes{
			TimeSlot: slot.TimeSlot,
			Status:   slot.Status,
		}
	}

	return timeSlotsDTO
}