package mappers

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/avaliability/dao"
)

func GetDailyAvailabilityCompetitorIDDAOtoDTO(dailyAvailabilityDAO *dao.GetDailyAvailabilityByIDDAORes) *models.GetDailyAvailabilityByIDDTORes {
	dailyAvailabilityDTO := &models.GetDailyAvailabilityByIDDTORes{
		Day:       dailyAvailabilityDAO.Day,
		TimeSlots: getDailyTimeSlotAvailabilityCompetitorIDDAOtoDTO(dailyAvailabilityDAO.TimeSlots),
	}

	return dailyAvailabilityDTO
}

func getDailyTimeSlotAvailabilityCompetitorIDDAOtoDTO(timeSlotsDAO []*dao.GetDailyTimeSlotByIDDAORes) []*models.GetDailyTimeSlotByIDDTORes {
	timeSlotsDTO := make([]*models.GetDailyTimeSlotByIDDTORes, len(timeSlotsDAO))

	for i, ts := range timeSlotsDAO {
		timeSlotsDTO[i] = &models.GetDailyTimeSlotByIDDTORes{
			TimeSlot: ts.TimeSlot,
			Status:   ts.Status,
		}
	}

	return timeSlotsDTO
}
