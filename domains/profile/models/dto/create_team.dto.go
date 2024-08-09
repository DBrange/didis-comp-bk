package dto

type CreateTeamDTOReq struct {
	Name         string   `json:"name"`
	TotalMembers int      `json:"total_members"`
	Image        *string   `json:"image"`
	AverageScore *float32  `json:"average_score"`
	Admins       []string `json:"admins"`
}
