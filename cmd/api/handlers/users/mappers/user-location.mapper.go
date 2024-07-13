package mappers

import (
	api_dto "github.com/DBrange/didis-comp-bk/cmd/api/handlers/users/dto"
	location_dto "github.com/DBrange/didis-comp-bk/domains/location/models/dto"
	user_dto "github.com/DBrange/didis-comp-bk/domains/user/models/dto"
)

func UserAndLocation(user *user_dto.GetUserByIDDTORes, location *location_dto.GetLocationByIDDTORes) *api_dto.GetUserByIDDTORes {
	completeUser := &api_dto.GetUserByIDDTORes{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Username:    user.Username,
		Email:       user.Email,
		Phone:       user.Phone,
		Birthdate:   user.Birthdate,
		Image:       user.Image,
		Active:      user.Active,
		Genre:       user.Genre,
		Role:        user.Role,
		AccessLevel: user.AccessLevel,
		Location:    location,
		PaymentID:   user.PaymentID,
		ScheduleID:  user.ScheduleID,
	}

	return completeUser
}
