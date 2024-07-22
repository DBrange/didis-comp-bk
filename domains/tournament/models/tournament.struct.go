package models

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tournament struct {
	ID                  primitive.ObjectID    `bson:"_id"`
	Name                string                `bson:"name"`
	FinishDate          *time.Time            `bson:"finish_date"`
	StartDate           *time.Time            `bson:"start_date"`
	Points              *int                  `bson:"points"`
	TotalPrize          float64               `bson:"total_prize"`
	TotalCompetitors    int                   `bson:"total_competitors"`
	MaxCapacity         int                   `bson:"max_capacity"`
	AverageScore        *float32              `bson:"average_score"`
	Genre               models.GENRE          `bson:"genre"`
	Sport               models.SPORT          `bson:"sport"`
	Surface             models.TENNIS_SURFACE `bson:"surface"`
	LocationID          primitive.ObjectID    `bson:"location_id"`
	OrganizerID         primitive.ObjectID    `bson:"organizer_id"`
	LeagueID            *primitive.ObjectID   `bson:"league_id"`
	DoubleEliminationID *primitive.ObjectID   `bson:"double_elimination_id"`
	Rounds              []primitive.ObjectID  `bson:"rounds"`
	Matches             []primitive.ObjectID  `bson:"matches"`
	Pots                []primitive.ObjectID  `bson:"pots"`
	Groups              []primitive.ObjectID  `bson:"groups"`
	CreatedAt           time.Time             `bson:"created_at"`
	UpdatedAt           time.Time             `bson:"updated_at"`
	DeletedAt           *time.Time            `bson:"deleted_at,omitempty"`
}
