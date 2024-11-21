package dao

import "go.mongodb.org/mongo-driver/bson/primitive"

type GetDayTimeSlotDAORes struct {
	ID                  *primitive.ObjectID               `bson:"_id"`
	DailyAvailabilities []*GetDailyAvailabilityByIDDAORes `bson:"daily_availabilities"`
}
