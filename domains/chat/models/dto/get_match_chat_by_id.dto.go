package dto

import "github.com/DBrange/didis-comp-bk/cmd/api/models"

type GetMatchChatByIDDTORes struct {
	ID                 string                               `json:"_id"`
	AvailabilityStatus models.CHAT_AVAILABILITY_STATUS      `json:"availability_status"`
	MatchID            string                               `json:"match_id"`
	Users              []*GetMatchChatByIDUserdDTORes       `json:"users"`
	Competitors        []*GetMatchChatByIDCompetitordDTORes `json:"competitors"`
}

type GetMatchChatByIDCompetitordDTORes struct {
	ID                 string                          `json:"_id"`
	AvailabilityStatus models.CHAT_AVAILABILITY_STATUS `json:"availability_status"`
	Users              []*GetMatchChatByIDUserdDTORes  `json:"users"`
}

type GetMatchChatByIDUserdDTORes struct {
	ID        string `json:"_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Image     string `json:"image"`
}
