package mappers

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/guest_competitor/dao"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

func CreateGuestCompetitorDTOtoDAO(guestCompetitorDTO *dto.CreateGuestCompetitorDTOReq, convert utils.ConvertToObjectIDFunc) (*dao.CreateGuestCompetitorDAOReq, error) {
	guestUserOID, err := convert(guestCompetitorDTO.GuestUserID)
	if err != nil {
		return nil, err
	}
	competitorOID, err := convert(guestCompetitorDTO.CompetitorID)
	if err != nil {
		return nil, err
	}

	guestCompetitorDAO := &dao.CreateGuestCompetitorDAOReq{
		GuestUserID:  *guestUserOID,
		CompetitorID: *competitorOID,
	}

	return guestCompetitorDAO, nil
}
