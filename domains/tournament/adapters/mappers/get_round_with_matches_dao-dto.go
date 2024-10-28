package mappers

import (
	"github.com/DBrange/didis-comp-bk/domains/repository/models/round/dao"
	"github.com/DBrange/didis-comp-bk/domains/tournament/models/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MapRoundWithMatchesDAOToDTO(dao *dao.GetRoundWithMatchesDAORes) *dto.GetRoundWithMatchesDTORes {
	if dao == nil {
		return nil
	}

	dto := &dto.GetRoundWithMatchesDTORes{
		ID:         mapObjectIDToString(dao.ID),
		Round:      dao.Round,
		TotalPrize: dao.TotalPrize,
		Points:     dao.Points,
		Matches:    mapMatchesDAOToDTO(dao.Matches),
	}
	return dto
}

func mapMatchesDAOToDTO(daoMatches []*dao.GetRoundWithMatchesMatchDAORes) []*dto.GetRoundWithMatchesMatchDTORes {
	if daoMatches == nil {
		return nil
	}

	dtoMatches := make([]*dto.GetRoundWithMatchesMatchDTORes, len(daoMatches))
	for i, match := range daoMatches {
		dtoMatches[i] = &dto.GetRoundWithMatchesMatchDTORes{
			ID:             mapObjectIDToString(match.ID),
			Date:           match.Date,
			Result:         match.Result,
			Position:       match.Position,
			PositionWinner: match.PositionWinner,
			Competitors:    mapCompetitorsDAOToDTO(match.Competitors),
		}
	}
	return dtoMatches
}

func mapCompetitorsDAOToDTO(daoCompetitors []*dao.GetRoundWithMatchesCompetitorDAORes) []*dto.GetRoundWithMatchesCompetitorDTORes {
	if daoCompetitors == nil {
		return nil
	}

	dtoCompetitors := make([]*dto.GetRoundWithMatchesCompetitorDTORes, len(daoCompetitors))
	for i, competitor := range daoCompetitors {
		dtoCompetitors[i] = &dto.GetRoundWithMatchesCompetitorDTORes{
			ID:              mapObjectIDToString(competitor.ID),
			CurrentPosition: competitor.CurrentPosition,
			Position:        competitor.Position,
			Users:           mapUsersDAOToDTO(competitor.Users),
			GuestUsers:      mapUsersDAOToDTO(competitor.GuestUsers),
		}
	}
	return dtoCompetitors
}

func mapUsersDAOToDTO(daoUsers []*dao.GetRoundWithMatchesUserDAORes) []*dto.GetRoundWithMatchesUserDTORes {
	if daoUsers == nil {
		return nil
	}

	dtoUsers := make([]*dto.GetRoundWithMatchesUserDTORes, len(daoUsers))
	for i, user := range daoUsers {
		dtoUsers[i] = &dto.GetRoundWithMatchesUserDTORes{
			ID:        mapObjectIDToString(user.ID),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Image:     user.Image,
		}
	}
	return dtoUsers
}

func mapObjectIDToString(id *primitive.ObjectID) *string {
	if id == nil {
		return nil
	}
	str := id.Hex()
	return &str
}
