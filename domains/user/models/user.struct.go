package models

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
)

type User struct {
	ID          string         `json:"id" bson:"_id"`
	FirstName   string         `json:"first_name" bson:"first_name"`
	LastName    string         `json:"last_name" bson:"last_name"`
	Username    *string         `json:"username" bson:"username"`
	Age         *int8           `json:"age" bson:"age"`
	Password    *string         `json:"password" bson:"password"`
	Email       string         `json:"email" bson:"email"`
	Phone       *string         `json:"phone" bson:"phone"`
	Image       *string         `json:"image" bson:"image"`
	Active      bool           `json:"active" bson:"active"`
	AccessLevel *int16          `json:"access_level" bson:"access_level"`
	Genre       []models.GENRE `json:"genre" bson:"genre"`
	Role        []models.ROLE  `json:"role" bson:"role"`
	LocationID  *string         `json:"location_id" bson:"location_id"`
	ScheduleID  *string         `json:"schedule_id" bson:"schedule_id"`
	PaymentID   *string         `json:"payment_id" bson:"payment_id"`
	CreatedAt   time.Time      `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" bson:"updated_at"`
	DeletedAt   *time.Time     `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}
