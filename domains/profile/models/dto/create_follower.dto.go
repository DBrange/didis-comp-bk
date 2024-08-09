package dto

type CreateFollowerDTOReq struct {
	From         string `json:"from"`
	ToUser       *string `json:"to_user"`
	ToCompetitor *string `json:"to_competitor"`
}
