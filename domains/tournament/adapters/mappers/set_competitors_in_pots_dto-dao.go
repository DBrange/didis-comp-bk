package mappers

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/pot/dao"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SetCompetitorsInPotsDTOtoDAO(potDTOs []*dto.SetPotCompetitorDTOReq, tournamentID string, convert utils.ConvertToObjectIDFunc) ([]*dao.SetPotCompetitorDAOReq, *primitive.ObjectID, error) {
	potDAOs := make([]*dao.SetPotCompetitorDAOReq, len(potDTOs))

	for i, potDTO := range potDTOs {
		competitorOIDs, err := utils.ConvertToObjectIDs(&potDTO.Competitors, convert)
		if err != nil {
			return nil,nil, err
		}

		potOID, err := convert(potDTO.PotID)
		if err != nil {
			return nil,nil, err
		}

		potDAOs[i] = &dao.SetPotCompetitorDAOReq{
			PotID:       potOID,
			Competitors: *competitorOIDs,
		}

	}

	tournamentOID, err := convert(tournamentID)
	if err != nil {
		return nil,nil, err
	}

	return potDAOs, tournamentOID, nil

}
