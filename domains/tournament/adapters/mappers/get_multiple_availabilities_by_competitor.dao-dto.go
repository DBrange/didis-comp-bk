package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/avaliability/dao"
)

func GetMultipleAvailabilitiesByCompetitor(availabilityDAOs [][]*dao.GetDailyAvailabilityByIDDAORes) [][]*dto.GetDailyAvailabilityByIDDTORes {
	availabilityDTOs := make([][]*dto.GetDailyAvailabilityByIDDTORes, len(availabilityDAOs))

	for i, availabilityDAO := range availabilityDAOs {
		availabilityDTOs[i] = getDalilyAvailabilityAvailabilities(availabilityDAO)
	}

	return availabilityDTOs
}

func getDalilyAvailabilityAvailabilities(availabilityDAOs []*dao.GetDailyAvailabilityByIDDAORes) []*dto.GetDailyAvailabilityByIDDTORes {
	dailyAvailabilityDTOs := make([]*dto.GetDailyAvailabilityByIDDTORes, len(availabilityDAOs))

	for i, dailyAvailabilityDAO := range availabilityDAOs {
		dailyAvailabilityDTOs[i] = &dto.GetDailyAvailabilityByIDDTORes{
			Day:       dailyAvailabilityDAO.Day,
			TimeSlots: getDalilyTymeSlotAvailabilityAvailabilities(dailyAvailabilityDAO.TimeSlots),
		}
	}

	return dailyAvailabilityDTOs
}

func getDalilyTymeSlotAvailabilityAvailabilities(timeSlotDAOs []*dao.GetDailyTimeSlotByIDDAORes) []*dto.GetDailyTimeSlotByIDDTORes {
	timeSlotDTOs := make([]*dto.GetDailyTimeSlotByIDDTORes, len(timeSlotDAOs))

	for i, timeSlotDAO := range timeSlotDAOs {
		timeSlotDTOs[i] = &dto.GetDailyTimeSlotByIDDTORes{
			TimeSlot: timeSlotDAO.TimeSlot,
			Status:   timeSlotDAO.Status,
		}
	}

	return timeSlotDTOs
}
