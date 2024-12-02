package dao

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateTournamentDAOReq struct {
	Name                string                     `bson:"name"`
	FinishDate          *time.Time                 `bson:"finish_date"`
	StartDate           *time.Time                 `bson:"start_date"`
	Points              *int                       `bson:"points"`
	Image              *string                       `bson:"image"`
	TotalPrize          float64                    `bson:"total_prize"`
	TotalCompetitors    int                        `bson:"total_competitors"`
	MaxCapacity         models.TOURNAMENT_CAPACITY `bson:"max_capacity"`
	AverageScore        float32                   `bson:"average_score"`
	Availability        TournamentAvailabilityDAO  `bson:"availability"`
	Genre               models.GENRE               `bson:"genre"`
	Sport               models.SPORT               `bson:"sport"`
	Surface             models.TENNIS_SURFACE      `bson:"surface"`
	CompetitorType      models.COMPETITOR_TYPE     `bson:"competitor_type"`
	LocationID          primitive.ObjectID         `bson:"location_id"`
	OrganizerID         primitive.ObjectID         `bson:"organizer_id"`
	CategoryID          *primitive.ObjectID        `bson:"category_id"`
	DoubleEliminationID *primitive.ObjectID        `bson:"double_elimination_id"`
	Rounds              []primitive.ObjectID       `bson:"rounds"`
	Matches             []primitive.ObjectID       `bson:"matches"`
	Pots                []primitive.ObjectID       `bson:"pots"`
	Groups              []primitive.ObjectID       `bson:"groups"`
	CreatedAt           time.Time                  `bson:"created_at"`
	UpdatedAt           time.Time                  `bson:"updated_at"`
	DeletedAt           *time.Time                 `bson:"deleted_at,omitempty"`
}

type TournamentAvailabilityDAO struct {
	AvailableCourts int `bson:"available_courts"`
	AverageHours    int `bson:"average_hours"`
}

func (u *CreateTournamentDAOReq) SetTimeStamp() {
	currentDate := time.Now().UTC()
	if u.CreatedAt.IsZero() {
		u.CreatedAt = currentDate
	}
	u.UpdatedAt = currentDate
}
