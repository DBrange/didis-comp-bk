package mappers

import (
	req_dto "github.com/DBrange/didis-comp-bk/cmd/api/handlers/users/dto"
	user_dto "github.com/DBrange/didis-comp-bk/domains/user/models/dto"
)

func OnlyUpdateUser(user *req_dto.UpdateUserDTOReq) *user_dto.UpdateUserDTOReq {
	onlyUser := &user_dto.UpdateUserDTOReq{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Phone:     user.Phone,
		Image:     user.Image,
	}

	return onlyUser
}
