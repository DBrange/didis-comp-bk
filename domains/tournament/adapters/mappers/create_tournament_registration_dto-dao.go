package mappers

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/tournament_registration/dao"
)

func CreateTournamentRegistrationDTOtoDAO(tournamentRegistrationDTO *dto.CreateTournamentRegistrationDTOReq, convert utils.ConvertToObjectIDFunc) (*dao.CreateTournamentRegistrationDAOReq, error) {
	competitorOID, err := convert(tournamentRegistrationDTO.CompetitorID)
	if err != nil {
		return nil, err
	}
	tournamentOID, err := convert(tournamentRegistrationDTO.TournamentID)
	if err != nil {
		return nil, err
	}

	tournamentRegistrationDAO := &dao.CreateTournamentRegistrationDAOReq{
		CompetitorID: *competitorOID,
		TournamentID: *tournamentOID,
	}

	return tournamentRegistrationDAO, nil
}
