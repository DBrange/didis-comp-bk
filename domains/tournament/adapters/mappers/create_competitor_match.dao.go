package mappers

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/competitor_match/dao"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

func CreateCompetitorMatchDTOtoDAO(competitorsDTOs *dto.CreateCompetitorMatchDTOReq, convert utils.ConvertToObjectIDFunc) (*dao.CreateCompetitorMatchDAOReq, error) {
	matchOID, err := convert(competitorsDTOs.MatchID)
	if err != nil {
		return nil, err
	}

	competitorsDAOs := &dao.CreateCompetitorMatchDAOReq{
		CompetitorID: nil,
		Position:     competitorsDTOs.Position,
		MatchID:      matchOID,
	}

	return competitorsDAOs, nil
}
