package mappers

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateAvailabilityDTOtODAO(userID, competitorID, tournamentID *string, convert utils.ConvertToObjectIDFunc) (userOID, competitorOID, tournamentOID *primitive.ObjectID, err error) {
	if userID != nil {
		userOID, err = convert(*userID)
		if err != nil {
			return nil, nil, nil, err
		}
	}

	if competitorID != nil {
		competitorOID, err = convert(*competitorID)
		if err != nil {
			return nil, nil, nil, err
		}
	}

	if tournamentID != nil {
		tournamentOID, err = convert(*tournamentID)
		if err != nil {
			return nil, nil, nil, err
		}
	}

	return
}
