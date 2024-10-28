package mappers

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/avaliability/dao"
)

func GetDailyAvailabilityByIDDAOtoDTO(availabilityInfoDAO *dao.GetDailyAvailabilityByIDDAORes) *models.GetDailyAvailabilityByIDDTORes {
	availabilityInfoDTO := &models.GetDailyAvailabilityByIDDTORes{
		Day:       availabilityInfoDAO.Day,
		TimeSlots: getTimeSlotDTORes(availabilityInfoDAO.TimeSlots),
	}

	return availabilityInfoDTO
}

func getTimeSlotDTORes(timeSlotInfoDTO []*dao.GetDailyTimeSlotByIDDAORes) []*models.GetDailyTimeSlotByIDDTORes {
	var timeSlotInfoDAO []*models.GetDailyTimeSlotByIDDTORes

	for _, slot := range timeSlotInfoDTO {
		slotDTO := &models.GetDailyTimeSlotByIDDTORes{
			TimeSlot: slot.TimeSlot,
			Status:   slot.Status,
		}

		timeSlotInfoDAO = append(timeSlotInfoDAO, slotDTO)
	}

	return timeSlotInfoDAO
}
