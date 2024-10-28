package mappers

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/avaliability/dao"
)

func GetAvailabilityByTournamentIDDAOtoDTO(dailyAvailabilitiesDAO []*dao.GetDailyAvailabilityByIDDAORes) []*models.GetDailyAvailabilityByIDDTORes {
	dailyAvailabilitiesDTO := make([]*models.GetDailyAvailabilityByIDDTORes, len(dailyAvailabilitiesDAO))

	for i, daily := range dailyAvailabilitiesDAO {
		dailyAvailabilitiesDTO[i] = &models.GetDailyAvailabilityByIDDTORes{
			Day:       daily.Day,
			TimeSlots: GetAvailabilityDailyTimeSlotSliceDTOtoDAO(daily.TimeSlots),
		}
	}

	return dailyAvailabilitiesDTO
}
