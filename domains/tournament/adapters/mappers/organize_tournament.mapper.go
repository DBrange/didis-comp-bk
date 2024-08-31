package mappers

import (
	tournament_dto "github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

func OrganizeTournamentMapper(organizerTournamentDTO *tournament_dto.OrganizeTournamentDTOReq) (*tournament_dto.CreateTournamentDTOReq, *tournament_dto.GetDoubleElimInfoToFinaliseItDTORes, *tournament_dto.CreateLocationDTOReq, *string, string) {

	tournamentInfoDTO := &tournament_dto.CreateTournamentDTOReq{
		Name:           organizerTournamentDTO.Name,
		Points:         organizerTournamentDTO.Points,
		TotalPrize:     organizerTournamentDTO.TotalPrize,
		MaxCapacity:    organizerTournamentDTO.MaxCapacity,
		Genre:          organizerTournamentDTO.Genre,
		Sport:          organizerTournamentDTO.Sport,
		Surface:        *organizerTournamentDTO.Surface,
		CompetitorType: organizerTournamentDTO.CompetitorType,
	}

	var doubleElim *tournament_dto.GetDoubleElimInfoToFinaliseItDTORes
	if organizerTournamentDTO.DoubleElim != nil {
		doubleElim = &tournament_dto.GetDoubleElimInfoToFinaliseItDTORes{
			TotalPrize: organizerTournamentDTO.DoubleElim.TotalPrize,
			Points:     *organizerTournamentDTO.DoubleElim.Points,
		}
	}

	locationInfoDTO := &tournament_dto.CreateLocationDTOReq{
		State:   organizerTournamentDTO.Location.State,
		Country: organizerTournamentDTO.Location.Country,
		City:    organizerTournamentDTO.Location.City,
		Lat:     organizerTournamentDTO.Location.Lat,
		Long:    organizerTournamentDTO.Location.Long,
	}

	return tournamentInfoDTO, doubleElim, locationInfoDTO, organizerTournamentDTO.CategoryID, organizerTournamentDTO.OrganizerID
}
