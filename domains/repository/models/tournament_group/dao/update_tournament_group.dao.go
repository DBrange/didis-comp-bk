package dao

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateTournamentGroupDAOReq struct {
	Competitors *[]primitive.ObjectID `bson:"competitors,omitempty"`
	Matches     *[]primitive.ObjectID `bson:"matches,omitempty"`
	UpdatedAt   time.Time             `bson:"updated_at"`
}

func (u *UpdateTournamentGroupDAOReq) RenewUpdate() {
	u.UpdatedAt = time.Now().UTC()
}
