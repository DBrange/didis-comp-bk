package mappers

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddMultipleMatchesInTournamentGroupDTOtoDAO(groupID, tournamentID string, matchIDs []string, convert utils.ConvertToObjectIDFunc) (*primitive.ObjectID, *primitive.ObjectID, []*primitive.ObjectID, error) {
	groupOID, err := convert(groupID)
	if err != nil {
		return nil, nil, nil, err
	}
	tournamentOID, err := convert(tournamentID)
	if err != nil {
		return nil, nil, nil, err
	}

	matchOIDs, err := utils.ConvertToObjectIDs(&matchIDs, convert)
	if err != nil {
		return nil, nil, nil, err
	}

	return groupOID, tournamentOID, *matchOIDs, nil

}
