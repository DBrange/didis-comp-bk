package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/avaliability/dao"
)

func GetDailyAvailabilityByIDDAOtoDTO(availabilityInfoDAO *dao.GetDailyAvailabilityByIDDAORes) *dto.GetDailyAvailabilityByIDDTORes {
	availabilityInfoDTO := &dto.GetDailyAvailabilityByIDDTORes{
		Day:       availabilityInfoDAO.Day,
		TimeSlots: getTimeSlotDTORes(availabilityInfoDAO.TimeSlots),
	}

	return availabilityInfoDTO
}

func getTimeSlotDTORes(timeSlotInfoDTO []*dao.GetDailyTimeSlotByIDDAORes) []*dto.GetDailyTimeSlotByIDDTORes {
	var timeSlotInfoDAO []*dto.GetDailyTimeSlotByIDDTORes

	for _, slot := range timeSlotInfoDTO {
		slotDTO := &dto.GetDailyTimeSlotByIDDTORes{
			TimeSlot: slot.TimeSlot,
			Status:   slot.Status,
		}

		timeSlotInfoDAO = append(timeSlotInfoDAO, slotDTO)
	}

	return timeSlotInfoDAO
}
