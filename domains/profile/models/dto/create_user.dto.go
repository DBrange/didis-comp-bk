package dto

import (
	"time"

	"github.com/DBrange/didis-comp-bk/cmd/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateUserDTOReq struct {
	// ID          string         `json:"id"`
	FirstName  string               `json:"first_name"`
	LastName   string               `json:"last_name"`
	Username   *string              `json:"username"`
	Birthdate  *time.Time           `json:"birthdate"`
	Password   *string              `json:"password"`
	Email      string               `json:"email"`
	Phone      *string              `json:"phone"`
	Image      *string              `json:"image"`
	Active     bool                 `json:"active"`
	Genre      models.GENRE         `json:"genre" validate:"genre"`
	Roles      []primitive.ObjectID `json:"role"`
	LocationID *string              `json:"location_id"`
	PaymentID  *string              `json:"payment_id"`
}

// func (u *CreateProfileDTOReq) SetTimeStamp() {
// 	currentDate := time.Now().UTC()
// 	if u.CreatedAt.IsZero() {
// 		u.CreatedAt = currentDate
// 	}
// 	u.UpdatedAt = currentDate
// }

// func (u *CreateProfileDTOReq) RenewUpdate() {
// 	u.UpdatedAt = time.Now().UTC()
// }
