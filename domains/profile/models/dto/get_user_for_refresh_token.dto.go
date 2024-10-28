package dto

type GetUserForRefreshTokenDTO struct {
	ID        string   `json:"id"`
	Roles     []string `json:"roles"`
}