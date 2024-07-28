package dto

type LoginDTOReq struct {
	Username string `json:"username"`
	Password string  `json:"password" validate:"password"`
}
