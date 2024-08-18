package dto

import "time"

type GetCategoriesFromOrganizerDTORes struct {
	CategoryID        string                                       `json:"category_id"`
	Competitors       []GetCategoriesFromOrganizerCompetitorDTORes `json:"competitors"`
	TotalParticipants int32                                        `json:"total_participants"`
}

type GetCategoriesFromOrganizerCompetitorDTORes struct {
	CompetitorID        string                                 `json:"competitor_id"`
	CurrentPosition     *int                                   `json:"current_position"`
	RegisteredPositions []RegistedPositionDTORes               `json:"registered_positions"`
	Points              int                                    `json:"points"`
	Users               []GetCategoriesFromOrganizerUserDTORes `json:"users"`
}

type GetCategoriesFromOrganizerUserDTORes struct {
	UserID    string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Image     string `json:"image"`
}

type RegistedPositionDTORes struct {
	Date     time.Time `json:"date"`
	Position int       `json:"position"`
}
