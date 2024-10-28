package mappers

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/team/dao"
	"github.com/DBrange/didis-comp-bk/domains/category/models/dto"
)

func CreateTeamDTOtoDAO(teamDTO *dto.CreateTeamDTOReq, convert utils.ConvertToObjectIDFunc) (*dao.CreateTeamDAOReq, error) {

	teamDAO := &dao.CreateTeamDAOReq{
		Name:         teamDTO.Name,
		Image:        teamDTO.Image,
		TotalMembers: teamDTO.TotalMembers,
		AverageScore: *teamDTO.AverageScore,
	}
	
	if teamDTO.Admins != nil {
		adminsOID, err := utils.ConvertToObjectIDs(&teamDTO.Admins, convert)
		if err != nil {
			return nil, err
		}

		teamDAO.Admins = *adminsOID
	}

	return teamDAO,nil
}
