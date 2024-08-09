package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

func ConvertToObjectIDs(IDs *[]string, convert ConvertToObjectIDFunc) (*[]primitive.ObjectID, error) {
	OIDs := make([]primitive.ObjectID, len(*IDs))
	for i, id := range *IDs {
		OID, err := convert(id)
		if err != nil {
			return nil, err
		}
		OIDs[i] = *OID
	}

	return &OIDs, nil
}
