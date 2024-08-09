package dto

type GetCompetitorFollowedDTORes struct {
	ID    string                             `json:"_id"`
	Users []*GetUserCompetitorFollowedDTORes `json:"users"`
}

type GetUserCompetitorFollowedDTORes struct {
	ID        string `json:"_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Image     string `json:"image"`
}
