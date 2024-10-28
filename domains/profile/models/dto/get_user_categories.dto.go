package dto

type GetUserCategoriesDTO struct {
	Categories []*GetUserCategoriesCategoryDTO `json:"categories"`
}

type GetUserCategoriesCategoryDTO struct {
	ID             string                              `json:"id"`
	Name           string                              `json:"name"`
	CompetitorData *GetUserCategoriesCompetitorDataDTO `json:"competitor_data"`
	Organizer      *GetUserCategoriesOrganizerDTO      `json:"organizer"`
}

type GetUserCategoriesOrganizerDTO struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type GetUserCategoriesCompetitorDataDTO struct {
	Points          int                         `json:"points"`
	CurrentPosition int                         `json:"current_position"`
	Users           []*GetUserCategoriesUserDTO `json:"users"`
}

type GetUserCategoriesUserDTO struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
