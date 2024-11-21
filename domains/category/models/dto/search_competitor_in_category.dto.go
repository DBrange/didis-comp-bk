package dto

type GetCompetitorsOfCategoryDTORes struct {
	Total       int64                                       `json:"total"`
	Competitors []*GetCompetitorsOfCategoryCompetitorDTORes `json:"competitors"`
}
type GetCompetitorsOfCategoryCompetitorDTORes struct {
	ID                  string                                `json:"id"`
	CurrentPosition     *int                                  `json:"current_position"`
	RegisteredPositions []RegistedPositionDTORes              `json:"registered_positions"`
	Points              int                                   `json:"points"`
	Users               []*GetCompetitorsOfCategoryUserDTORes `json:"users"`
	GuestUsers          []*GetCompetitorsOfCategoryUserDTORes `json:"guest_users"`
}

type GetCompetitorsOfCategoryUserDTORes struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Image     string `json:"image"`
	Username     string `json:"username"`
}
