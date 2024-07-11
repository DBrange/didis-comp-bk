package mappers

import (
	req_dto "github.com/DBrange/didis-comp-bk/cmd/api/handlers/users/dto"
	location_dto "github.com/DBrange/didis-comp-bk/domains/location/models/dto"
)

func OnlyCreateLocation(user *req_dto.CreateUserDTOReq) *location_dto.CreateLocationDTOReq {
	onlyLocation := &location_dto.CreateLocationDTOReq{
		State:   user.Location.State,
		Country: user.Location.Country,
		City:    user.Location.City,
		Lat:     user.Location.Lat,
		Long:    user.Location.Long,
	}

	return onlyLocation
}
