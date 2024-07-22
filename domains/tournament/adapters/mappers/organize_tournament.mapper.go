package mappers

import (
	location_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/tournament"
	tournament_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/tournament/dao"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

func OrganizeTournamentMapper(tournamentInfoDTO *dto.OrganizeTournamentDTOReq) (*tournament_dao.CreateTournamentDAOReq, *location_dao.CreateLocationDAOReq, *tournament.OrganizeTournamentOptions, *string, string) {

	tournamentInfoDAO := &tournament_dao.CreateTournamentDAOReq{
		Name:             tournamentInfoDTO.Name,
		Points:           tournamentInfoDTO.Points,
		TotalPrize:       tournamentInfoDTO.TotalPrize,
		TotalCompetitors: tournamentInfoDTO.TotalCompetitors,
		MaxCapacity:      tournamentInfoDTO.MaxCapacity,
		Genre:            tournamentInfoDTO.Genre,
		Sport:            tournamentInfoDTO.Sport,
		Surface:          *tournamentInfoDTO.Surface,
	}

	locationInfoDAO := &location_dao.CreateLocationDAOReq{
		State:   tournamentInfoDTO.Location.State,
		Country: tournamentInfoDTO.Location.Country,
		City:    tournamentInfoDTO.Location.City,
		Lat:     tournamentInfoDTO.Location.Lat,
		Long:    tournamentInfoDTO.Location.Long,
	}

	organizeTournamentOptions := &tournament.OrganizeTournamentOptions{
		DoubleEliminationID: tournamentInfoDTO.DoubleElimination,
		Groups:              tournamentInfoDTO.Groups,
		Pots:                tournamentInfoDTO.Pots,
	}

	return tournamentInfoDAO, locationInfoDAO, organizeTournamentOptions, tournamentInfoDTO.LeagueID, tournamentInfoDTO.OrganizerID
}
