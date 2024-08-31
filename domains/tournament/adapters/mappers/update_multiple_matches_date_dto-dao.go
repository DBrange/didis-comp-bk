package mappers

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/match/dao"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

func UpdateMultipleMatchesDateDTOtoDAO(matchesDateDTO []*dto.MatchDateDTOReq, convert utils.ConvertToObjectIDFunc) ([]*dao.MatchDateDAOReq, error) {
	matchesDateDAO := make([]*dao.MatchDateDAOReq, len(matchesDateDTO))

	for i, mdDAO := range matchesDateDTO {
		OID, err := convert(mdDAO.ID)
		if err != nil {
			return nil, err
		}

		matchesDateDAO[i] = &dao.MatchDateDAOReq{
			ID:   OID,
			Date: mdDAO.Date,
		}
	}

	return matchesDateDAO, nil
}
