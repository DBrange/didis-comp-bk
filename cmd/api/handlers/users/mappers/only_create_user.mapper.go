package mappers

import (
	req_dto "github.com/DBrange/didis-comp-bk/cmd/api/handlers/users/dto"
	user_dto "github.com/DBrange/didis-comp-bk/domains/user/models/dto"
)

func OnlyCreateUser(user *req_dto.RegisterUserDTOReq) *user_dto.CreateUserDTOReq {
	onlyUser := &user_dto.CreateUserDTOReq{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Birthdate: user.Birthdate,
		Email:     user.Email,
		Phone:     user.Phone,
		Genre:     user.Genre,
		Password:  user.Password,
		Image:     user.Image,
	}

	return onlyUser
}
