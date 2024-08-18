package dto

type AddCompetitorsToTournamentGroupsDTOReq struct {
	GroupID     string   `json:"group_id" validate:"required"`
	Competitors []string `json:"competitors" validate:"required,min=1,dive,required"`
}
