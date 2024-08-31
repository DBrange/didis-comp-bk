package dao

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateDoubleEliminationDAOReq struct {
	Matches    *[]*primitive.ObjectID `bson:"matches,omitempty"`
	Rounds     *[]*primitive.ObjectID `bson:"rounds,omitempty"`
	TotalPrize *float64               `bson:"total_prize,omitempty"`
	Points     *int                   `bson:"points,omitempty"`
	UpdatedAt  time.Time              `bson:"updated_at"`
}

func (u *UpdateDoubleEliminationDAOReq) RenewUpdate() {
	u.UpdatedAt = time.Now().UTC()
}
