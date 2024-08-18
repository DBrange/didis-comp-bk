package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/round/dao"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetRoundsWithCompetitorsDAOtoDTO(roundDAOs []*dao.GetRoundWithCompetitorsDAORes) []*dto.GetRoundWithCompetitorsDTORes {
	roundDTOs := make([]*dto.GetRoundWithCompetitorsDTORes, len(roundDAOs))

	for i, roundDAO := range roundDAOs {
		roundDTOs[i] = &dto.GetRoundWithCompetitorsDTORes{
			ID:            roundDAO.ID.Hex(),
			CompetitorIDs: getRoundsWithCompetitorsCompetitorsDAOtoDTO(roundDAO.CompetitorIDs),
			TotalPrize:    roundDAO.TotalPrize,
			Points:        roundDAO.Points,
		}
	}

	return roundDTOs
}

func getRoundsWithCompetitorsCompetitorsDAOtoDTO(competitorOIDs []*primitive.ObjectID) []string {
	competitorIDs := make([]string, len(competitorOIDs))

	for i, cID := range competitorOIDs {
		competitorIDs[i] = cID.Hex()
	}

	return competitorIDs
}
