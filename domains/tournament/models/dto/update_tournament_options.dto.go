package dto

import "github.com/DBrange/didis-comp-bk/domains/repository/models/common"

type UpdateTournamentOptionsDTOReq struct {
	DoubleEliminationID  *string   `json:"double_elimination_id,omitempty"`
	Pots                 *[]string `json:"pots,omitempty"`
	Groups               *[]string `json:"groups,omitempty"`
	Matches              *[]string `json:"matches,omitempty"`
	Rounds               *[]string `json:"rounds,omitempty"`
	common.UpdateBaseDAO `json:",inline"`
}
