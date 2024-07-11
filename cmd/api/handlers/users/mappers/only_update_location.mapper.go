package mappers

import (
	req_dto "github.com/DBrange/didis-comp-bk/cmd/api/handlers/users/dto"
	location_dto "github.com/DBrange/didis-comp-bk/domains/location/models/dto"
)

func OnlyUpdateLocation(user *req_dto.UpdateUserDTOReq) *location_dto.UpdateLocationDTOReq {
	onlyLocation := &location_dto.UpdateLocationDTOReq{
		ID:      user.Location.ID,
		State:   user.Location.State,
		Country: user.Location.Country,
		City:    user.Location.City,
		Lat:     user.Location.Lat,
		Long:    user.Location.Long,
	}

	return onlyLocation
}
