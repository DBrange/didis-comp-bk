package dao

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
)

type CreateUserDAO struct {
	FirstName   string         `bson:"first_name"`
	LastName    string         `bson:"last_name"`
	Username    *string        `bson:"username"`
	Birthdate   *time.Time     `bson:"birthdate"`
	Password    *string        `bson:"password"`
	Email       string         `bson:"email"`
	Phone       *string        `bson:"phone"`
	Image       *string        `bson:"image"`
	Active      bool           `bson:"active"`
	AccessLevel *int16         `bson:"access_level"`
	Genre       []models.GENRE `bson:"genre"`
	Role        []models.ROLE  `bson:"role"`
	LocationID  *string        `bson:"location_id"`
	ScheduleID  *string        `bson:"schedule_id"`
	PaymentID   *string        `bson:"payment_id"`
	CreatedAt   time.Time      `bson:"created_at"`
	UpdatedAt   time.Time      `bson:"updated_at"`
	DeletedAt   *time.Time     `bson:"deleted_at,omitempty"`
}

func (u *CreateUserDAO) SetTimeStamp() {
	currentDate := time.Now().UTC()
	if u.CreatedAt.IsZero() {
		u.CreatedAt = currentDate
	}
	u.UpdatedAt = currentDate
}

func (u *CreateUserDAO) RenewUpdate() {
	u.UpdatedAt = time.Now().UTC()
}
