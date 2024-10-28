package dto

type GetCompetitorFollowedDTORes struct {
	ID              string                             `json:"id"`
	CurrentPosition *int                               `json:"current_position"`
	Users           []*GetUserCompetitorFollowedDTORes `json:"users"`
	GuestUsers      []*GetUserCompetitorFollowedDTORes `json:"guest_users"`
}

type GetUserCompetitorFollowedDTORes struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Image     string `json:"image"`
	Username  string `json:"username"`
}
