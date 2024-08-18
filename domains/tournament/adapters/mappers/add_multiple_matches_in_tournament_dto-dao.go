package mappers

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddMultipleMatchesInTournamentDTOtoDAO( tournamentID string, matchIDs []string, convert utils.ConvertToObjectIDFunc) ( *primitive.ObjectID, []*primitive.ObjectID, error) {
	tournamentOID, err := convert(tournamentID)
	if err != nil {
		return nil, nil, err
	}

	matchOIDs, err := utils.ConvertToObjectIDs(&matchIDs, convert)
	if err != nil {
		return nil, nil, err
	}

	return tournamentOID, *matchOIDs, nil

}
