package mappers

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/avaliability/dao"
)

func GetMultipleAvailabilitiesByCompetitor(availabilityDAOs [][]*dao.GetDailyAvailabilityByIDDAORes) [][]*models.GetDailyAvailabilityByIDDTORes {
	availabilityDTOs := make([][]*models.GetDailyAvailabilityByIDDTORes, len(availabilityDAOs))

	for i, availabilityDAO := range availabilityDAOs {
		availabilityDTOs[i] = getDalilyAvailabilityAvailabilities(availabilityDAO)
	}

	return availabilityDTOs
}

func getDalilyAvailabilityAvailabilities(availabilityDAOs []*dao.GetDailyAvailabilityByIDDAORes) []*models.GetDailyAvailabilityByIDDTORes {
	dailyAvailabilityDTOs := make([]*models.GetDailyAvailabilityByIDDTORes, len(availabilityDAOs))

	for i, dailyAvailabilityDAO := range availabilityDAOs {
		dailyAvailabilityDTOs[i] = &models.GetDailyAvailabilityByIDDTORes{
			Day:       dailyAvailabilityDAO.Day,
			TimeSlots: getDalilyTymeSlotAvailabilityAvailabilities(dailyAvailabilityDAO.TimeSlots),
		}
	}

	return dailyAvailabilityDTOs
}

func getDalilyTymeSlotAvailabilityAvailabilities(timeSlotDAOs []*dao.GetDailyTimeSlotByIDDAORes) []*models.GetDailyTimeSlotByIDDTORes {
	timeSlotDTOs := make([]*models.GetDailyTimeSlotByIDDTORes, len(timeSlotDAOs))

	for i, timeSlotDAO := range timeSlotDAOs {
		timeSlotDTOs[i] = &models.GetDailyTimeSlotByIDDTORes{
			TimeSlot: timeSlotDAO.TimeSlot,
			Status:   timeSlotDAO.Status,
		}
	}

	return timeSlotDTOs
}
