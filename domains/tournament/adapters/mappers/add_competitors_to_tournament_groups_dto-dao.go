package mappers

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/tournament_group/dao"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddCompetitorsToTournamentGroupsDTOtoDAO(groupDTOs []*dto.AddCompetitorsToTournamentGroupsDTOReq, tournamentID string, convert utils.ConvertToObjectIDFunc) ([]*dao.AddCompetitorsToTournamentGroupsDAOReq, *primitive.ObjectID, error) {
	groupDAOs := make([]*dao.AddCompetitorsToTournamentGroupsDAOReq, len(groupDTOs))

	tournamentOID, err := convert(tournamentID)
	if err != nil {
		return nil, nil, err
	}

	for i, groupDTO := range groupDTOs {
		competitorsOIDs, err := utils.ConvertToObjectIDs(&groupDTO.Competitors, convert)
		if err != nil {
			return nil, nil, err
		}

		groupOID, err := convert(groupDTO.GroupID)
		if err != nil {
			return nil, nil, err
		}

		groupDAOs[i] = &dao.AddCompetitorsToTournamentGroupsDAOReq{
			GroupID:     groupOID,
			Competitors: *competitorsOIDs,
		}
	}

	return groupDAOs, tournamentOID, nil
}
