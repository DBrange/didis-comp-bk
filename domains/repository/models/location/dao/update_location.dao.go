package dao

import "time"

type UpdateLocationDAOReq struct {
	ID        string     `bson:"_id,omitempty"`
	State     *string    `bson:"state,omitempty"`
	Country   *string    `bson:"country,omitempty"`
	City      *string    `bson:"city,omitempty"`
	Lat       *string    `bson:"lat,omitempty"`
	Long      *string    `bson:"long,omitempty"`
	UpdatedAt *time.Time `bson:"updated_at,omitempty"`
}

func (u *UpdateLocationDAOReq) RenewUpdate() {
	currentTime := time.Now().UTC()
	u.UpdatedAt = &currentTime
}
