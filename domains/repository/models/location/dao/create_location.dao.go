package dao

import "time"

type CreateLocationDAOReq struct {
	State     *string     `bson:"state"`
	Country   *string     `bson:"country"`
	City      *string     `bson:"city"`
	Lat       *string     `bson:"lat"`
	Long      *string     `bson:"long"`
	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
}

func (u *CreateLocationDAOReq) SetTimeStamp() {
	currentDate := time.Now()
	if u.CreatedAt.IsZero() {
		u.CreatedAt = currentDate
	}
	u.UpdatedAt = currentDate
}

func (u *CreateLocationDAOReq) RenewUpdate() {
	u.UpdatedAt = time.Now()
}
