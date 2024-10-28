package mappers

import (

	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/competitor_match/dao"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateMultipleCompetitorMatchesDTOtoDAO(competitorMatchDTOs []*dto.UpdateCompetitorMatchDTOReq, convert utils.ConvertToObjectIDFunc) ([]*dao.UpdateCompetitorMatchDAOReq, error) {
	competitorMatchDAOs := make([]*dao.UpdateCompetitorMatchDAOReq, len(competitorMatchDTOs))

	for i, competitorMatchDTO := range competitorMatchDTOs {
		matchOID, err := convert(competitorMatchDTO.MatchID)
		if err != nil {
			return nil, err
		}

		var competitorOID *primitive.ObjectID
		if competitorMatchDTO.CompetitorID != nil{
			competitorOID, err = convert(*competitorMatchDTO.CompetitorID)
			if err != nil {
				return nil, err
			}
		}

		competitorMatchDAOs[i] = &dao.UpdateCompetitorMatchDAOReq{
			MatchID:      *matchOID,
			CompetitorID: competitorOID,
			Position:     competitorMatchDTO.Position,
		}

	}

	return competitorMatchDAOs, nil
}
