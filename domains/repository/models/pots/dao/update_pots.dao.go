package dao

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdatePotsDAOReq struct {
	Competitors []primitive.ObjectID `bson:"competitors"`
	UpdatedAt   time.Time            `bson:"updated_at"`
}

func (u *UpdatePotsDAOReq) RenewUpdate() {
	u.UpdatedAt = time.Now().UTC()
}
