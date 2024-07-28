package mappers

// func ModifyProfileAvailabilityMapper(availabilityInfoDTO *dto.ModifyProfileAvailabilityDTOReq) *dao.UpdateAvailabilityDAOReq {
// 	dailyAvailabilityDAO := updateDailyAvailabilityDAOReq(availabilityInfoDTO.DailyAvailabilities)

// 	availabilityInfoDAO := &dao.UpdateAvailabilityDAOReq{
// 		DailyAvailabilities: dailyAvailabilityDAO,
// 	}

// 	return availabilityInfoDAO
// }

// func ModifyProfileDailyAvailabilityMapper(dailyAvailabilityDTO *dto.ModifyProfileDailyAvailabilityDTOReq) *dao.UpdateDailyAvailabilityDAOReq {
// 	dailyAvailabilityDAO := &dao.UpdateDailyAvailabilityDAOReq{
// 		Day:       dailyAvailabilityDTO.Day,
// 		TimeSlots: updateTimeSlotDAOReqMapper(dailyAvailabilityDTO.TimeSlots),
// 	}

// 	return dailyAvailabilityDAO
// }

// func updateDailyAvailabilityDAOReq(dailyAvailabilityDTO []*dto.ModifyProfileDailyAvailabilityDTOReq) []*dao.UpdateDailyAvailabilityDAOReq {
// 	var dailyAvailabilityDAO []*dao.UpdateDailyAvailabilityDAOReq

// 	for _, day := range dailyAvailabilityDTO {
// 		dailyDAO := &dao.UpdateDailyAvailabilityDAOReq{
// 			Day:       day.Day,
// 			TimeSlots: updateTimeSlotDAOReqMapper(day.TimeSlots),
// 		}

// 		dailyAvailabilityDAO = append(dailyAvailabilityDAO, dailyDAO)
// 	}

// 	return dailyAvailabilityDAO
// }

// func updateTimeSlotDAOReqMapper(timeSlotInfoDTO []*dto.ModifyProfileTimeSlotDTOReq) []*dao.UpdateTimeSlotDAOReq {
// 	var timeSlotInfoDAO []*dao.UpdateTimeSlotDAOReq

// 	for _, slot := range timeSlotInfoDTO {
// 		slotDAO := &dao.UpdateTimeSlotDAOReq{
// 			TimeSlot: slot.TimeSlot,
// 			Status:   slot.Status,
// 		}

// 		timeSlotInfoDAO = append(timeSlotInfoDAO, slotDAO)
// 	}

// 	return timeSlotInfoDAO
// }
