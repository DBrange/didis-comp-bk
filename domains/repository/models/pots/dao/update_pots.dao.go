package dao

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdatePotDAOReq struct {
	Competitors []primitive.ObjectID `bson:"competitors"`
	UpdatedAt   time.Time            `bson:"updated_at"`
}

func (u *UpdatePotDAOReq) RenewUpdate() {
	u.UpdatedAt = time.Now().UTC()
}
