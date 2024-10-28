package dto

type CreateGuestCompetitorDTOReq struct {
	GuestUserID  string `json:"guest_user_id"`
	CompetitorID string `json:"competitor_id"`
}
