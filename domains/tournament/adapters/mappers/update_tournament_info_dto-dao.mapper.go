package mappers

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/tournament/dao"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateTournamentInfoDTOtoDAO(tournamentDTO *dto.UpdateTournamentInfoDTOReq, tournamentID string, convert utils.ConvertToObjectIDFunc) (*dao.UpdateTournamentInfoDAOReq, *primitive.ObjectID,error) {
	tournamentOID, err := convert(tournamentID)
	if err != nil{
		return nil,nil, err
	}
	
	tournamentDAO := &dao.UpdateTournamentInfoDAOReq{
		Name:         tournamentDTO.Name,
		Points:       tournamentDTO.Points,
		TotalPrize:   tournamentDTO.TotalPrize,
		AverageScore: tournamentDTO.AverageScore,
		Surface:      tournamentDTO.Surface,
		StartDate:    tournamentDTO.StartDate,
		FinishDate:   tournamentDTO.FinishDate,
	}

	return tournamentDAO, tournamentOID, nil
}
