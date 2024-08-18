package dto

type GetPositionsBracketMatchDTORes struct {
	ID            string                                      `json:"_id"`
	PositionMatch int                                         `json:"position_match"`
	Competitors   []*GetPositionsBracketMatchCompetitorDTORes `json:"competitors"`
}

type GetPositionsBracketMatchCompetitorDTORes struct {
	ID              string `json:"_id"`
	Position        int    `json:"position"`
	CurrentPosition *int   `json:"current_position"`
}

type GetPositionsMatch struct {
	ID            string
	PositionMatch int
}
