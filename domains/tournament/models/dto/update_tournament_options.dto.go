package dto

import "github.com/DBrange/didis-comp-bk/domains/repository/models/common"

type UpdateTournamentOptionsDTOReq struct {
	DoubleEliminationID  *string   `json:"double_elimination_id,omitempty"`
	Rounds               *[]string `json:"rounds,omitempty"`
	Pots                 *[]string `json:"pots,omitempty"`
	Groups               *[]string `json:"groups,omitempty"`
	common.UpdateBaseDAO `json:",inline"`
}
