package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/round/dao"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
)

func GetRoundGroupsDAOtoDTO(dao *dao.GetRoundGroupsDAORes) *dto.GetRoundGroupsDTORes {
	if dao == nil {
		return nil
	}

	dto := &dto.GetRoundGroupsDTORes{
		ID:              dao.ID.Hex(),
		Round:           dao.Round,
		TotalPrize:      dao.TotalPrize,
		Points:          dao.Points,
		TotalClassified: dao.TotalClassified,
		BestThird:       dao.BestThird,
		Groups:          GetRoundGroupDAOtoDTO(dao.Groups),
	}

	return dto
}

func GetRoundGroupDAOtoDTO(daoGroups []*dao.GetRoundGroupDAORes) []*dto.GetRoundGroupDTORes {
	if daoGroups == nil {
		return nil
	}

	dtoMatches := make([]*dto.GetRoundGroupDTORes, len(daoGroups))
	for i, group := range daoGroups {

		dtoMatches[i] = &dto.GetRoundGroupDTORes{
			ID:          group.ID.Hex(),
			Position:    group.Position,
			Matches:     mapMatchesDAOToDTO(group.Matches),
			Competitors: mapCompetitorsWithStatsDAOToDTO(group.Competitors),
		}
	}
	return dtoMatches
}

func mapCompetitorsWithStatsDAOToDTO(daoCompetitors []*dao.GetRoundGroupCompetitorWithStatsDAORes) []*dto.GetRoundGroupCompetitorWithStatsDTORes {
	if daoCompetitors == nil {
		return nil
	}

	dtoCompetitors := make([]*dto.GetRoundGroupCompetitorWithStatsDTORes, len(daoCompetitors))
	for i, competitor := range daoCompetitors {
		dtoCompetitors[i] = &dto.GetRoundGroupCompetitorWithStatsDTORes{
			ID:              *mapObjectIDToString(competitor.ID),
			CurrentPosition: competitor.CurrentPosition,
			Position:        competitor.Position,
			Stats:           *mapStatsDAOToDTO(&competitor.Stats),
			Users:           mapUsersDAOToDTO(competitor.Users),
			GuestUsers:      mapUsersDAOToDTO(competitor.GuestUsers),
		}
	}
	return dtoCompetitors
}
func mapStatsDAOToDTO(statsDAO *dao.TournamentGroupCompetitorStatsDAOReq) *dto.TournamentGroupCompetitorStatsDTOReq {
	if statsDAO == nil {
		return nil
	}

	statsDTO := &dto.TournamentGroupCompetitorStatsDTOReq{
		MatchesPlayed:   statsDAO.MatchesPlayed,
		MatchesLost:     statsDAO.MatchesLost,
		MatchesWon:      statsDAO.MatchesWon,
		SetsWon:         statsDAO.SetsWon,
		SetsLost:        statsDAO.SetsLost,
		GamesWon:        statsDAO.GamesWon,
		GamesLost:       statsDAO.GamesLost,
		LastFiveMatches: statsDAO.LastFiveMatches,
	}

	return statsDTO
}
