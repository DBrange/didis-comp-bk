package dao

import (
	"time"

	location_dao "github.com/DBrange/didis-comp-bk/domains/repository/models/location/dao"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetTournamentsFromCategoryDAORes struct {
	ID           primitive.ObjectID                  `bson:"_id"`
	Name         string                              `bson:"name"`
	Points       *int                                `bson:"points"`
	Location     *location_dao.GetLocationByIDDAORes `bson:"location"`
	TotalPrize   float64                             `bson:"total_prize"`
	AverageScore int                                 `bson:"average_score"`
	StartDate    *time.Time                          `bson:"start_date"`
	FinishtDate  *time.Time                          `bson:"finish_date"`
}
