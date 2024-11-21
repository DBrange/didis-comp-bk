package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/match/dao"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

func GetMatchByIDDAOToDTO(daoMatches *dao.GetMatchDAORes) *dto.GetMatchDTORes {
	if daoMatches == nil {
		return nil
	}

	dtoMatches := &dto.GetMatchDTORes{
		ID:             daoMatches.ID.Hex(),
		Date:           daoMatches.Date,
		Result:         daoMatches.Result,
		Position:       daoMatches.Position,
		PositionWinner: daoMatches.PositionWinner,
		Sport:          daoMatches.Sport,
		Competitors:    mapCompetitorsDAOToDTO(daoMatches.Competitors),
		Tournament:     GetMatchByIDTournamentDAOToDTO(daoMatches.Tournament),
		Round:          GetMatchByIDRoundDAOToDTO(daoMatches.Round),
	}
	return dtoMatches
}

func GetMatchByIDTournamentDAOToDTO(daoTournament *dao.GetMatchTorurnamentDAORes) *dto.GetMatchTorurnamentDTORes {
	return &dto.GetMatchTorurnamentDTORes{
		ID:   daoTournament.ID.Hex(),
		Name: daoTournament.Name,
	}
}
func GetMatchByIDRoundDAOToDTO(daoTournament *dao.GetMatchRoundDAORes) *dto.GetMatchRoundDTORes {
	return &dto.GetMatchRoundDTORes{
		ID:    daoTournament.ID.Hex(),
		Round: daoTournament.Round,
	}
}
