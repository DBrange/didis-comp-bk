package mappers

import (
	availability_dto "github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	availability_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/avaliability/dao"
)

func GetProfileAvailabilityInfoByIDMapper(availabilityInfoDAO *availability_dao.GetDailyAvailabilityInfoByIDDAORes) *availability_dto.GetProfileDailyAvailabilityInfoByIDDTORes {
	availabilityInfoDTO := &availability_dto.GetProfileDailyAvailabilityInfoByIDDTORes{
		Day:       availabilityInfoDAO.Day,
		TimeSlots: getTimeSlotDTOResMapper(availabilityInfoDAO.TimeSlots),
	}

	return availabilityInfoDTO
}

func getTimeSlotDTOResMapper(timeSlotInfoDTO []*availability_dao.GetDailyTimeSlotInfoByIDDAORes) []*availability_dto.GetProfileDailyTimeSlotInfoByIDDTORes {
	var timeSlotInfoDAO []*availability_dto.GetProfileDailyTimeSlotInfoByIDDTORes

	for _, slot := range timeSlotInfoDTO {
		slotDTO := &availability_dto.GetProfileDailyTimeSlotInfoByIDDTORes{
			TimeSlot: slot.TimeSlot,
			Status:   slot.Status,
		}

		timeSlotInfoDAO = append(timeSlotInfoDAO, slotDTO)
	}

	return timeSlotInfoDAO
}
