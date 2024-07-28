package dto

type CreateGuestCompetitorDTOReq struct {
	GuestUserID  string `json:"guest_competitor_id"`
	CompetitorID string `json:"competitor_id"`
}
