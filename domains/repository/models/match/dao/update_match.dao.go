package dao

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateMatchDAOReq struct {
	Result   *string             `bson:"result,omitempty"`
	Winner   *primitive.ObjectID `bson:"winner,omitempty"`
	Position int                 `bson:"position,omitempty"`
	// Votes        map[string]string  `bson:"votes,omitempty"`
	UpdatedAt time.Time `bson:"updated_at,omitempty"`
}

func (u *UpdateMatchDAOReq) RenewUpdate() {
	u.UpdatedAt = time.Now().UTC()
}
