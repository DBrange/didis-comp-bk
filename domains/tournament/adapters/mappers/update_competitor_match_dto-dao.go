package mappers

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/competitor_match/dao"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateCompetitorMatchDTOtoDAO(competitorMatchDTO *dto.UpdateCompetitorMatchDTOReq, matchID string, convert utils.ConvertToObjectIDFunc) (*dao.UpdateCompetitorMatchDAOReq, *primitive.ObjectID, error) {
	matchOID, err := convert(matchID)
	if err != nil {
		return nil, nil, err
	}

	competitorOID, err := convert(*competitorMatchDTO.CompetitorID)
	if err != nil {
		return nil, nil, err
	}

	competitorMatchDAO := &dao.UpdateCompetitorMatchDAOReq{
		CompetitorID: competitorOID,
		Position:     competitorMatchDTO.Position,
	}

	return competitorMatchDAO, matchOID, nil
}