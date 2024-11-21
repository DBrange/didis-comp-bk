package dto

type GetUserPrimatyInfoDTORes struct {
	User        *GetUserPrimaryDataDTORes `json:"user"`
	Followers   int                       `json:"followers"`
	IsFollowing bool                      `json:"is_following"`
	OrganizerID *string                    `json:"organizer_id"`
}
