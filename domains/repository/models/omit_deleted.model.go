package models

import "go.mongodb.org/mongo-driver/bson"

// Omit all schemes with deleted_at or deleted_at with content
func OmitDeleted() bson.M {
	return bson.M{
		"$or": []bson.M{
			{"deleted_at": bson.M{"$exists": false}},
			{"deleted_at": nil},
		},
	}
}

