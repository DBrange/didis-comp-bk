package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/avaliability/dao"
)

func UpdateAvailabilityDTOtoDAO(availabilityInfoDTO *dto.UpdateAvailabilityDTOReq) *dao.UpdateAvailabilityDAOReq {
	dailyAvailabilityDAO := updateDailyAvailabilityDAOReq(availabilityInfoDTO.DailyAvailabilities)

	availabilityInfoDAO := &dao.UpdateAvailabilityDAOReq{
		DailyAvailabilities: dailyAvailabilityDAO,
	}

	return availabilityInfoDAO
}

func UpdateDailyAvailabilityDTOtoDAO(dailyAvailabilityDTO *dto.UpdateDailyAvailabilityDTOReq) *dao.UpdateDailyAvailabilityDAOReq {
	dailyAvailabilityDAO := &dao.UpdateDailyAvailabilityDAOReq{
		Day:       dailyAvailabilityDTO.Day,
		TimeSlots: updateTimeSlotDAOReq(dailyAvailabilityDTO.TimeSlots),
	}

	return dailyAvailabilityDAO
}

func updateDailyAvailabilityDAOReq(dailyAvailabilityDTO []*dto.UpdateDailyAvailabilityDTOReq) []*dao.UpdateDailyAvailabilityDAOReq {
	var dailyAvailabilityDAO []*dao.UpdateDailyAvailabilityDAOReq

	for _, day := range dailyAvailabilityDTO {
		dailyDAO := &dao.UpdateDailyAvailabilityDAOReq{
			Day:       day.Day,
			TimeSlots: updateTimeSlotDAOReq(day.TimeSlots),
		}

		dailyAvailabilityDAO = append(dailyAvailabilityDAO, dailyDAO)
	}

	return dailyAvailabilityDAO
}

func updateTimeSlotDAOReq(timeSlotInfoDTO []*dto.UpdateTimeSlotDTOReq) []*dao.UpdateTimeSlotDAOReq {
	var timeSlotInfoDAO []*dao.UpdateTimeSlotDAOReq

	for _, slot := range timeSlotInfoDTO {
		slotDAO := &dao.UpdateTimeSlotDAOReq{
			TimeSlot: slot.TimeSlot,
			Status:   slot.Status,
		}

		timeSlotInfoDAO = append(timeSlotInfoDAO, slotDAO)
	}

	return timeSlotInfoDAO
}
