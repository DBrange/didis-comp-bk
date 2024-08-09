package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/avaliability/dao"
)

func GetDailyAvailabilityCompetitorIDDAOtoDTO(dailyAvailabilityDAO *dao.GetDailyAvailabilityByIDDAORes) *dto.GetDailyAvailabilityByIDDTORes {
	dailyAvailabilityDTO := &dto.GetDailyAvailabilityByIDDTORes{
		Day:       dailyAvailabilityDAO.Day,
		TimeSlots: getDailyTimeSlotAvailabilityCompetitorIDDAOtoDTO(dailyAvailabilityDAO.TimeSlots),
	}

	return dailyAvailabilityDTO
}

func getDailyTimeSlotAvailabilityCompetitorIDDAOtoDTO(timeSlotsDAO []*dao.GetDailyTimeSlotByIDDAORes) []*dto.GetDailyTimeSlotByIDDTORes {
	timeSlotsDTO := make([]*dto.GetDailyTimeSlotByIDDTORes, len(timeSlotsDAO))

	for i, ts := range timeSlotsDAO {
		timeSlotsDTO[i] = &dto.GetDailyTimeSlotByIDDTORes{
			TimeSlot: ts.TimeSlot,
			Status:   ts.Status,
		}
	}

	return timeSlotsDTO
}
